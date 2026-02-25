# Lesson 3: Consistent Hashing (The Ring)

## The Problem: Scaling the Cache
You built a Redis Rate Limiter. Great.
But now you have **100 Million Users**. One Redis instance cannot hold all the keys.
You need **10 Redis Instances**.

**Naive Approach (Modulo):**
`server_index = hash(user_id) % 10`
- User A -> Server 1
- User B -> Server 2

**The Disaster:**
Server 9 crashes. You now have **9 servers**.
The formula changes: `hash(user_id) % 9`.
**Result:** almost EVERY user moves to a different server.
Your cache hit rate drops to 0%. The database melts down.

## The Solution: Consistent Hashing
Instead of a line (0..9), imagine a **Ring** (0..360 degrees).
1.  **Place Servers on the Ring:** Hash their IPs to pick a spot.
2.  **Place Users on the Ring:** Hash their IDs.
3.  **Assignment:** Walk **clockwise** from the User to find the first Server.

**Why it works:**
If Server 9 dies, only the users "behind" Server 9 (between Server 8 and 9) move to Server 10 (or 0).
Everyone else stays put.
**Only 1/N keys move.**

---

## Assignment: The Ring Simulator

We will build a Go library that implements Consistent Hashing.

**Struct:**
```go
type ConsistentHash struct {
    keys       []int          // Sorted hash codes of servers
    hashMap    map[int]string // Map hash -> Server Name
    replicas   int            // Virtual nodes (to balance the ring)
}
```

**Task:**
1.  Create `Academy/Week-02-Distributed-Patterns/consistent-hash/ring.go`.
2.  Implement `Add(node string)`:
    - Hash the node + replica ID (e.g., "Server1-0", "Server1-1").
    - Store code in `keys` and `hashMap`.
    - Sort `keys` (Crucial for the ring walk).
3.  Implement `Get(key string) string`:
    - Hash the key.
    - Binary Search (`sort.Search`) in `keys` to find the first node >= hash.
    - Return that node.

**Go.**
