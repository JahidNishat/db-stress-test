# Lesson 5: Optimistic Locking (The "Version" Strategy)

## The Problem with `FOR UPDATE` (Pessimistic Locking)
`FOR UPDATE` is safe, but it is **SLOW**.
It puts a physical lock on the row.
If User A is reading the row, User B waits.
It turns your database into a single-file line (Serial).
This kills performance at scale.

## The Solution: Optimistic Locking
Instead of locking the row ("I suspect someone will rob me"), we assume everything will be fine ("I am optimistic").
But we verify at the end.

**How it works:**
1.  Add a `version` column to the table.
2.  **Read:** `SELECT val, version FROM counter WHERE id=1` (e.g., val=0, version=1).
3.  **Calculate:** `new_val = val + 1` (in Go).
4.  **Write:**
    ```sql
    UPDATE counter 
    SET val = $1, version = version + 1 
    WHERE id = 1 AND version = $2
    ```
    (Where `$2` is the OLD version we read: 1).

**The Magic:**
If someone else updated the row while we were thinking... the `version` in the DB is now 2.
Our query says `WHERE version = 1`.
The DB finds **0 rows**.
The update fails.
We know we failed. We can **Retry**.

---

## Assignment: Implement Optimistic Locking

1.  **Modify Schema:** Add `version INT DEFAULT 1` to `counter` table in `schema.sql`. (You might need to drop table or alter it).
2.  **Modify Go Code:**
    - Change `SELECT` to get `version` too.
    - Change `UPDATE` to check `WHERE version = old_version`.
    - **Check the result:** `res, _ := tx.Exec(...)`.
    - `rowsAffected, _ := res.RowsAffected()`.
    - If `rowsAffected == 0`, it means **CONFLICT**.
    - **Crucial:** If conflict, you must **Loop and Retry** until you succeed.

**Task:**
Rewrite `worker` to use this loop.
1.  Loop forever.
2.  Start Tx. Read.
3.  Try Update.
4.  If Success -> Commit & Break.
5.  If Fail (0 rows updated) -> Rollback & Continue Loop.

**Go.**
