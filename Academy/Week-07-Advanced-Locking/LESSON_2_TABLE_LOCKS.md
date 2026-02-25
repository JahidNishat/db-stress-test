# Week 7, Lesson 2: Table Locks (The Nuclear Option)

## The Analogy: Renovating the Building

### 1. Row Locking (`FOR UPDATE`)
- Locking one book.
- People can still use the rest of the library.
- Fast, high-concurrency.

### 2. Table Locking (`LOCK TABLE`)
- Locking the **whole building**.
- Nobody can enter. Nobody can read. Nobody can write.
- This is what happens when you run `ALTER TABLE` or `DROP TABLE`.

---

## The Danger of `ACCESS EXCLUSIVE`
In production, this is the #1 cause of "Service Outages."
If a long-running migration (like adding a column to a 1G table) starts, it will block **every single query** hitting that table.

## Levels of Table Locks
Postgres actually has many levels (8 in total!), but these are the big ones:

1. **Row Share:** (e.g., `SELECT ... FOR SHARE`) - Blocks nothing except Table Locks.
2. **Row Exclusive:** (e.g., `INSERT`, `UPDATE`) - Blocks Table Locks.
3. **Access Exclusive:** (e.g., `ALTER TABLE`) - Blocks **everything**.

---

## Senior Tip: The "Safe Migration"
When adding a column, always check if it requires a "Rewrite" of the table. In modern Postgres (11+), adding a column with a `DEFAULT NULL` is instant and safe. Adding it with a complex `DEFAULT value` can lock the table for minutes!
