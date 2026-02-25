# Senior Backend Engineer Assessment (Months 1-2)

**Role:** Senior Go Backend Engineer
**Interviewer:** Principal Architect
**Time Limit:** Take your time, but answer with precision.

---

## Part 1: Go Internals & Concurrency

### Q1. The "Memory Leak"
You are reviewing a Junior Dev's code. They have a function that reads a 1GB file into memory, processes it, and returns a small `*Result` struct.
They notice that even after the function returns, the RAM usage stays high.
**Code Snippet:**
```go
func ProcessFile() *Result {
    data, _ := os.ReadFile("large_file.txt") // 1GB slice
    // ... processing ...
    return &Result{Summary: string(data[:100])} // Returning a slice of the array
}
```
**Question:** Why is the memory not being released? How would you fix it?

### Q2. The "Slow" Channel
You have a producer generating 1,000 events/sec. You have a consumer that takes 10ms to process an event.
You use a buffered channel `make(chan Event, 100)`.
**Question:** What happens to the producer after 1 second? How would you architect this system to handle "bursts" without dropping data?

### Q3. Escape Analysis
If I pass a `*http.Request` to a function `Process(r *http.Request)`, and inside that function I launch `go func() { Use(r) }()`, does `r` stay on the Stack or move to the Heap? Why?

---

## Part 2: Distributed Systems

### Q4. The Load Balancer Logic
We have 3 backend servers. We use **Consistent Hashing** to route users.
**Scenario:** Server B crashes.
1. What happens to the users who were mapped to Server B?
2. What happens to the users who were mapped to Server A?
3. Why is this better than `hash(userID) % 3`?

### Q5. Rate Limiter (Redis)
You implemented a "Fixed Window" rate limiter in Redis (100 req/min).
**Scenario:** A user sends 100 requests at 12:00:59 and another 100 requests at 12:01:01.
**Question:** Does the Fixed Window limiter allow this? Is this a problem? How would you fix it using a different algorithm?

---

## Part 3: Database Internals (Postgres)

### Q6. The "Slow" Pagination
A dashboard query `SELECT * FROM orders ORDER BY created_at LIMIT 10 OFFSET 500000` is timing out.
**Question:** Explain *exactly* why this is slow (what does the DB engine do?). Propose a fix that makes it O(1) or O(log N).

### Q7. Indexing Strategy
You have a query: `SELECT * FROM users WHERE last_name = 'Smith' AND age > 25`.
Which index is better?
A. `(age, last_name)`
B. `(last_name, age)`
**Question:** Choose one and explain why (mention B-Tree structure).

### Q8. The "Missing" Update
Two transactions run at the same time:
**Tx1:** `UPDATE products SET stock = stock - 1 WHERE id = 1;` (Stock was 10)
**Tx2:** `UPDATE products SET stock = stock - 1 WHERE id = 1;` (Stock was 10)
**Question:** In `Read Committed` isolation (default), what is the final stock? 8 or 9? Explain how locking handles this.

---

## Part 4: Advanced Locking

### Q9. The "Report Generation" Problem
You need to run a heavy report on the `orders` table (takes 30 seconds). You don't want to block new orders from being inserted (`INSERT`), but you want to make sure no existing orders are modified (`UPDATE`) while you read them.
**Question:** Which lock mode do you use? `FOR UPDATE`, `FOR SHARE`, or `LOCK TABLE`?

### Q10. The Queue Architecture
You are building a Job Queue using Postgres. You have 50 worker processes.
They all run: `SELECT * FROM jobs WHERE status = 'pending' LIMIT 1 FOR UPDATE`.
**Question:** Performance is terrible. Workers are freezing. Why? What 2 words can you add to the SQL query to fix this instantly?

---

**[End of Assessment]**
**Please write your answers below each question.**

---

**[ANSWERS BEGIN]**

---

## Part 1: Go Internals & Concurrency

### Answer Q1:
Here the file is not closed, we should close it with defer... also we are ignoring the error... if we close without ignoring the error data can be nil pointer when there occurs error. this is the reason of memory leak
```go
func ProcessFile() *Result {
    data, err := os.ReadFile("large_file.txt") // 1GB slice
    if err != nil {
        panic(err)
    }
    defer data.Close()
    // ... processing ...
    return &Result{Summary: string(data[:100])} // Returning a slice of the array
}
```

### Answer Q2:
We can here use wait group to control the data... We have to add every wait group before launching the go func and defer done and have to wait until all the wait group done...

### Answer Q3:
It will move to the heap cause here it calls the closer with pointer value... so go runtime will asume the value can be used later so it moves r to the heap

---

## Part 2: Distributed Systems

### Answer Q4:
Let, we are created a hash ring, and everything is counterwise A->B->C->A, so if B crashes then the users of B goes to C, The user who were mapped to A will be unchanged they will still go to A, 1/3 of the users will move only. Its better than hash(userID)%3 because in this case about 66% user will move if a server changes, but for consistant hash it will be 33%.

### Answer Q5:
Yes a user can request 200 req in that time interval, the Fixed window allows this, and its one of the disadvantages of it. To fix this we can use, `Token Bucket` algorithm, it allows burst and individual user req and works perfect as a rate limiter. In this algorithm, we use a filling rate and max token that a user can access, for every req we decrese one limit until hits zero and every second it fills users limit by its rate filling speed.

---

## Part 3: Database Internals (Postgres)

### Q6. The "Slow" Pagination

### Answer Q6:
Using offset the db runs seq scan that is go by each id continously which is slow. We can use cursors to be fast fetching and cause in cursor it knows the exact location from where to begin and its index helps it also.

### Answer Q7:
last_name, age cause the query is about first last_name and then age, but theres an interesting things in postgresql, when it plan to exute the query it checks the index and plan query in optimum way, so both query works. But heres we select all the fields so it will find the rows by the index but still use heap scan cause it needs all the fields, so it will be still slow, if we don't need all the fields we can fetch the age, last_name for better performance also we can use include index to get exact field.

### Answer Q8:
The postgres default isolation is read committed, it prevents dirty read, but doesn't prevents no-repeated read and phantom read, so by its nature both txn will read 10 and decrese by one, so final result will be 9 altough it has to be 8. This locks reads the value in the txn during time and if another txn update the values it won't be count that. it will work on its old value.

---

## Part 4: Advanced Locking

### Answer Q9:
I will use For share lock, cause it will allows the use to modify the rows, for update will will block the read and write so i will choose for share thus it blocks write not read. And forget about lock table cause it will lock the entire table and nobody can insert new orders.

### Answer Q10:
We can use skip locked, because postgres will manage a queue and it will give worker the next available process:
`SELECT * FROM jobs WHERE status = 'pending' LIMIT 1 FOR UPDATE SKIP LOCKED;`

---
**[END OF ANSWERS]**

---