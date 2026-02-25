# Lesson 2: The Conveyor Belt (Channels)

## The Problem with Mutexes
In the last lesson, we used a **Mutex**.
A Mutex is like a **Toilet Key**.
Only one person can use the resource at a time. Everyone else stands in line waiting.
If you have 1000 goroutines, 999 of them are **PAUSED** (Blocked) waiting for the key.
This is safe, but can be slow if everyone is fighting for the same key.

## The Go Way: Channels
Go's motto:
> "Do not communicate by sharing memory; share memory by communicating."

Instead of 1000 chefs fighting over **one counter**, imagine a **Conveyor Belt**.
1.  The 2000 goroutines do not touch the counter.
2.  Instead, they throw the number "1" onto a conveyor belt (Channel).
3.  **One Single Chef** stands at the end of the belt. He picks up the numbers and adds them up.

Because only **one** person touches the final counter, **YOU DO NOT NEED A LOCK.**
There is no race condition because there is no competition.

---

## Assignment: Refactor to Channels

We are going to change `main.go` completely.

**Requirements:**
1.  Delete the `Mutex`. Keep the `WaitGroup`.
2.  Create a channel that accepts integers: `ch := make(chan int)`
3.  **The Workers:**
    - Change the 2000 goroutines.
    - Instead of `counter++`, they should send the value `1` into the channel: `ch <- 1`.
    - They still need `wg.Done()`.
4.  **The Collector (The Monitor):**
    - You need a separate goroutine (or loop) that reads from the channel.
    - `val := <- ch`
    - `counter += val`
5.  **The Tricky Part (Deadlock):**
    - If you wait for `wg.Wait()` *before* you close the channel, the collector might hang forever waiting for more data.
    - You might need a separate goroutine to close the channel when work is done.

**Challenge:** Try to implement this. It is harder than the Mutex.
If you get a "Deadlock" error, do not panic. It means everyone is waiting and no one is working.

**Hint Structure:**
```go
func main() {
    ch := make(chan int)
    var wg sync.WaitGroup
    
    // 1. Start the Workers (Producers)
    // ... loop 2000 times ...
    // ... go func() { ch <- 1; wg.Done() } ...

    // 2. Start the Closer (Manager)
    // You need a goroutine that waits for WG, then closes the channel.
    go func() {
        wg.Wait()
        close(ch)
    }()

    // 3. The Main Thread (Consumer)
    // Loop over the channel until it is closed.
    counter := 0
    for n := range ch {
        counter += n
    }
    
    fmt.Println(counter)
}
```
**Go!**
