# Week 3: Database Internals - The Pagination Trap

## The Problem: `OFFSET` is $O(N)$

When you run `SELECT * FROM users LIMIT 10 OFFSET 900000`, Postgres doesn't just jump to row 900,000. It must:
1.  Read all 900,000 rows from disk/buffer.
2.  Discard them.
3.  Take the next 10.

As your table grows to millions of rows, page 1 is fast, but page 10,000 is slow and heavy on I/O.

## The Solution: Keyset Pagination (Cursors)

Instead of saying "skip 900,000 rows", we say "give me the next 10 rows after the last ID I saw".
`SELECT * FROM users WHERE id > 900000 ORDER BY id ASC LIMIT 10`

If `id` is indexed, this is a B-Tree lookup + 10 rows. It's $O(1)$ or $O(\log N)$ regardless of how deep you are in the data.

---

## Task: Proving it with Go

We will build a script to:
1.  Seed a table with 1,000,000 rows.
2.  Benchmark `OFFSET` at different depths (100, 10,000, 500,000, 900,000).
3.  Benchmark `Keyset` (Cursor) at the same depths.
4.  Compare the execution time.

### Implementation Checklist
- [ ] Create `schema.sql` for a large dataset.
- [ ] Write `seeder.go` to insert 1M rows efficiently (use `COPY` or multi-insert).
- [ ] Write `benchmark.go` to run the queries and measure time.
