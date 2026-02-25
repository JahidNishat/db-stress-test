# Week 8, Lesson 2: Composite Indexes & The Leftmost Rule

## The Concept
A composite index `(A, B)` is sorted by `A` first, and then by `B`.

## The Rule (Leftmost Prefix)
Postgres can use the index if your query filters on:
1. The first column (`A`).
2. The first AND second column (`A` and `B`).

It **cannot** use the index if you only filter on the second column (`B`), because `B` is not sorted globally (it's only sorted *within* each `A`).

## The Experiment
- **Index:** `(name, age)`
- **Query:** `WHERE name = 'X'` -> **Index Scan** (Good).
- **Query:** `WHERE name = 'X' AND age = 20` -> **Index Scan** (Good).
- **Query:** `WHERE age = 20` -> **Seq Scan** (Bad).

## Senior Tip: "WHERE" Order Doesn't Matter
Postgres is smart. 
`WHERE age = 20 AND name = 'X'`
is exactly the same as
`WHERE name = 'X' AND age = 20`.
The Query Planner reorders them to match the index.
