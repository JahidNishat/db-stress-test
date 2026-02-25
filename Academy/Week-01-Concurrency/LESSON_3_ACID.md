# Lesson 3: The ACID Trip

## Why Databases Don't Just "Work"

You use Postgres. You write `UPDATE users SET balance = balance - 10 WHERE id = 1`.
It feels like magic. But under the hood, Postgres is doing exactly what we just did with Mutexes, but on a massive scale.

## The 4 Pillars (ACID)

1.  **A - Atomicity (All or Nothing)**
    *   *The Bank Transfer:* I take $10 from you. I give $10 to Bob.
    *   If the power fails after I take your money but *before* I give it to Bob...
    *   **Atomicity** ensures the $10 reappears in your account. The transaction is "Rolled Back".
    *   *Real world:* `BEGIN` ... `COMMIT` / `ROLLBACK`.

2.  **C - Consistency (Rules are Rules)**
    *   *The Constraint:* Account balance cannot be negative.
    *   If a transaction tries to make it -$5, the database says "NO" and cancels the whole thing.

3.  **I - Isolation (The Invisible Shield)**
    *   *This is what we studied today.*
    *   If I am reading your balance, and you are updating your balance... what do I see?
    *   **Read Committed:** I see the old value until you commit.
    *   **Serializable:** It acts like we are in a single-file line (Mutex).

4.  **D - Durability (Written in Stone)**
    *   If the server crashes *0.1ms after* I say "Success", the data MUST be there when it reboots.
    *   This is why databases are slow. They have to write to the physical hard drive (WAL - Write Ahead Log) before confirming.

---

## Assignment: The Bank Heist (Simulation)

We are going to simulate a bank with **Race Conditions**.

**The Setup:**
1.  Create `Academy/Week-01-Concurrency/bank.go`.
2.  Struct `Account` with `balance int` and a `Mutex`.
3.  Function `Transfer(from, to *Account, amount int)`.
4.  **The Bug:**
    *   Lock `from`. Check balance. Deduct money. Unlock `from`.
    *   ... (Artificial Delay) ...
    *   Lock `to`. Add money. Unlock `to`.
5.  **The Attack:**
    *   Run this concurrently.
    *   What happens if I check the total money in the system *during* the delay?
    *   Money will "disappear" for a moment. This violates **Atomicity** (viewed from outside) and **Consistency**.

**Goal:**
Write the broken code first. Show me that money disappears temporarily.
Then we will fix it.

**Structure:**
```go
type Account struct {
    sync.Mutex
    balance int
}

func Transfer(from, to *Account, amount int) {
    // Phase 1: Withdraw
    from.Lock()
    from.balance -= amount
    from.Unlock()

    time.Sleep(1 * time.Millisecond) // The "Network Lag"

    // Phase 2: Deposit
    to.Lock()
    to.balance += amount
    to.Unlock()
}
```
Run a main function that continually prints the `Total System Balance` (Alice + Bob). It should flicker.
$$