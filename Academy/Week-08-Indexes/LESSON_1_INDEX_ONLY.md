# Week 8, Lesson 1: Index Only Scans (The Holy Grail)

## The Concept
Normally, an index is a **Pointer**.
1.  **Index Scan:** Postgres looks in the index to find the Row ID (CTID).
2.  **Heap Fetch:** Postgres goes to the main table (Heap) to get the actual data.
3.  **Result:** Two disk reads (one for index, one for data).

## The Goal: "Index Only Scan"
If the index contains **ALL** the columns you asked for in your `SELECT`, Postgres doesn't need to visit the Heap. It just answers directly from the Index.
- **Result:** One disk read. 2x faster.

---

## The Experiment
1.  **Index:** `CREATE INDEX idx_age ON users (age);`
2.  **Query 1:** `SELECT age FROM users WHERE age = 25;`
    - The index has `age`. The query asks for `age`.
    - **Result:** Index Only Scan (Fast).

3.  **Query 2:** `SELECT name FROM users WHERE age = 25;`
    - The index has `age`, but NOT `name`.
    - **Result:** Index Scan + Heap Fetch (Slow).

## Senior Trick: `INCLUDE` (Covering Index)
To make Query 2 an "Index Only Scan," you can do:
`CREATE INDEX idx_age_include_name ON users (age) INCLUDE (name);`
- This stores `name` in the leaf nodes of the B-Tree without sorting by it.
- **Result:** Query 2 becomes an Index Only Scan!

---

## Reflection
- `SELECT *` breaks Index Only Scans because you almost never have an index with *every single column*.
- Indexes slow down `INSERT` / `DELETE` because you have to update the B-Tree every time.
