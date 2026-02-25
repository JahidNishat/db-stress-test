# Week 6: Isolation Levels & MVCC

## The Goal
In Week 5, we saw how to lock data. In Week 6, we learn how the database handles "Read vs Write" conflicts without locking everything.

## The 3 Database "Phenomena" (Bugs you want to avoid)

1.  **Dirty Read:** You read data that another transaction has changed but **not yet committed**. If they rollback, your data is fake!
2.  **Non-repeatable Read:** You read a row twice in one transaction, but the values change because someone else committed a change in between.
3.  **Phantom Read:** You query a range (e.g., `SELECT COUNT(*) WHERE age > 20`), and a new row is inserted by someone else, making your count wrong.

---

## Postgres Isolation Levels

Postgres has 4 levels, but effectively uses 3:
- **Read Committed (Default):** Prevents Dirty Reads.
- **Repeatable Read:** Prevents Dirty Reads and Non-repeatable Reads.
- **Serializable:** The gold standard. Prevents all phenomena (but is slow).

---

## Task 1: The "Invisible" Transaction
We are going to prove that by default, Postgres protects you from **Dirty Reads**.

1.  Open two terminals.
2.  **Terminal 1:** Start a transaction and change Alice's balance to $999,999. (Do NOT commit).
3.  **Terminal 2:** Try to read Alice's balance.

**Predict:** Will Terminal 2 see the $999,999 or the old $100?
