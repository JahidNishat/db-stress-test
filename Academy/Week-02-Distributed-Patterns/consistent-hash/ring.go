package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	c := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if c.hash == nil {
		c.hash = crc32.ChecksumIEEE
	}
	return c
}

func (m *Map) AddNode(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			virtualNodeName := key + "-" + strconv.Itoa(i)
			hash := int(m.hash([]byte(virtualNodeName)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}

func main() {
	// 1. Create Ring with 3 Replicas
	hashRing := New(3, nil)

	// 2. Add Servers
	hashRing.AddNode("ServerA", "ServerB", "ServerC")

	// 3. Check where users go
	users := []string{"User1", "User2", "User3", "User4", "User5", "User6"}
	fmt.Println("--- Initial Distribution ---")
	for _, u := range users {
		server := hashRing.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}

	// 4. Add a New Server (Scaling Up)
	fmt.Println("\n--- Adding ServerD ---")
	hashRing.AddNode("ServerD")

	// 5. Check again (Did everyone move? Or just a few?)
	for _, u := range users {
		server := hashRing.Get(u)
		fmt.Printf("%s => %s\n", u, server)
	}
}
