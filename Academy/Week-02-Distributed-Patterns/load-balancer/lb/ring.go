package main

import (
	"hash/crc32"
	"slices"
	"sort"
	"strconv"
	"sync"
)

type ConsistentHash struct {
	mu       sync.RWMutex
	hashMap  map[uint32]string
	keys     []uint32
	replicas int
}

func New(replicas int) *ConsistentHash {
	return &ConsistentHash{
		hashMap:  make(map[uint32]string),
		keys:     make([]uint32, 0),
		replicas: replicas,
	}
}

func (ch *ConsistentHash) Add(node string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	for i := 0; i < ch.replicas; i++ {
		virtualNodeName := node + "#" + strconv.Itoa(i)
		hash := crc32.ChecksumIEEE([]byte(virtualNodeName))
		ch.hashMap[hash] = node
		ch.keys = append(ch.keys, hash)
	}
	slices.Sort(ch.keys)
}

func (ch *ConsistentHash) Remove(node string) {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	for i := 0; i < ch.replicas; i++ {
		virtualNodeName := node + "#" + strconv.Itoa(i)
		hash := crc32.ChecksumIEEE([]byte(virtualNodeName))
		delete(ch.hashMap, hash)

		// Correctly remove from slice
		for j, v := range ch.keys {
			if v == hash {
				ch.keys = append(ch.keys[:j], ch.keys[j+1:]...)
				break
			}
		}
	}
}

func (ch *ConsistentHash) Get(key string) string {
	ch.mu.RLock()
	defer ch.mu.RUnlock()
	hash := crc32.ChecksumIEEE([]byte(key))
	idx := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i] >= hash
	})

	if len(ch.keys) == 0 {
		return "" // No servers available
	}

	if idx == len(ch.keys) {
		idx = 0
	}
	return ch.hashMap[ch.keys[idx]]
}

func (ch *ConsistentHash) Contains(node string) bool {
	ch.mu.RLock()
	defer ch.mu.RUnlock()
	for i := 0; i < ch.replicas; i++ {
		virtualNodeName := node + "#" + strconv.Itoa(i)
		hash := crc32.ChecksumIEEE([]byte(virtualNodeName))
		if _, ok := ch.hashMap[hash]; !ok {
			return false
		}
	}
	return true
}

func (ch *ConsistentHash) Nodes() []string {
	ch.mu.RLock()
	defer ch.mu.RUnlock()

	unique := map[string]bool{}
	for _, hash := range ch.hashMap {
		unique[hash] = true
	}

	var addrs []string
	for key := range unique {
		addrs = append(addrs, key)
	}
	return addrs
}
