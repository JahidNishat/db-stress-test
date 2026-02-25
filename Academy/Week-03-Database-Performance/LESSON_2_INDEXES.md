# Week 3, Lesson 2: The Magic of B-Trees & The Index Tax

## 1. The Power of Indexes
We proved that searching for a specific record in a 1,000,000 row table dropped from **~50ms** (Sequential Scan) to **~0.3ms** (Index Scan) once we added an index on the `email` column.

## 2. EXPLAIN ANALYZE
This is your primary tool for debugging slow queries.
- **Seq Scan:** The DB reads the whole table from disk. $O(N)$.
- **Index Scan:** The DB uses a B-Tree to find the pointer to the data. $O(\log N)$.
- **Parallel Seq Scan:** The DB uses multiple CPU cores to scan the table faster, but it's still heavy on I/O.

## 3. The "Index Tax" (Slow Writes)
We observed that adding 3 extra indexes significantly slowed down our `seedData` process.
- **Why?** Every `INSERT` now has to update 4 separate B-Tree structures.
- **Senior Rule:** Only index columns you actually query in your `WHERE` or `JOIN` clauses. Avoid indexing every column "just in case."

## 4. Wildcards
- `LIKE 'ABC%'`: **Uses Index.** The B-Tree can find the range starting with ABC.
- `LIKE '%ABC'`: **Cannot use Index.** The DB must scan every row to check the suffix.
