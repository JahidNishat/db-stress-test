# Lesson 4: Breaking the Deadlock (Law of Order)

## The Problem
You saw the Deadlock.
- Thread A locks Alice, wants Bob.
- Thread B locks Bob, wants Alice.
- Infinite wait.

## The Solution: Global Lock Ordering
If everyone agrees to **always pick up the fork on the left first**, or **always lock the Account with the smaller ID first**, deadlocks vanish.

**Scenario with Order (Alice ID=1, Bob ID=2):**
1.  **Thread A (Alice -> Bob):** Wants locks. Checks IDs. 1 < 2. Locks Alice. Then Locks Bob.
2.  **Thread B (Bob -> Alice):** Wants locks. Checks IDs. 1 < 2. **IT MUST LOCK ALICE FIRST.**
3.  Thread B tries to lock Alice. Alice is already locked by Thread A.
4.  Thread B waits *before* it touches Bob.
5.  Thread A is free to grab Bob (because Thread B is waiting at the door for Alice, not holding Bob).
6.  Thread A finishes. Thread B proceeds.

---

## Assignment: Fix `bank.go`

1.  Add `ID string` (or `int`) to the `Account` struct.
2.  In `main`, give them IDs: `alice.ID = "1"`, `bob.ID = "2"`.
3.  In `Transfer`, add logic to sort the locks:

```go
func Transfer(from, to *Account, amount int) {
    // Determine order
    var first, second *Account
    if from.ID < to.ID {
        first = from
        second = to
    } else {
        first = to
        second = from
    }

    // Lock in order
    first.Lock()
    second.Lock()

    // Move money
    from.balance -= amount
    to.balance += amount

    // Unlock (Order doesn't matter here, but reverse is polite)
    second.Unlock()
    first.Unlock()
}
```

**Task:** Apply this fix. Run the code.
It should run forever with NO "Fraud" and NO Deadlock.
