# Week 5, Lesson 2: Deadlocks & The Ordering Rule

## 1. What is a Deadlock?
A situation where two transactions are waiting for each other to release a lock.
- **T1:** Locks A, waits for B.
- **T2:** Locks B, waits for A.

Postgres detects this after a few seconds and kills one transaction with error code `40P01`.

## 2. The Solution: Consistent Ordering
Deadlocks are mathematically impossible if you always lock resources in the same order.
- **Bad:** One function locks `ID:1` then `ID:2`, another locks `ID:2` then `ID:1`.
- **Good:** Both functions always sort the IDs and lock the smaller ID first.

## 3. The `defer tx.Rollback()` Pattern
Always use `defer tx.Rollback()` after `db.Begin()`. 
- If `Commit()` is called, the rollback does nothing.
- If the code errors or panics, the rollback ensures the lock is released instantly.
