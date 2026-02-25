# Week 8, Lesson 3: GIN Indexes (JSON & Arrays)

## The Problem with B-Trees
B-Trees are great for scalar values (Int, String, Date).
But they fail for complex types like:
- **JSONB:** `{"tags": ["a", "b"]}`
- **Arrays:** `[1, 2, 3]`
- **Full Text Search:** `to_tsvector('hello world')`

## The Solution: GIN (Generalized Inverted Index)
GIN stores a mapping of `Key -> List of Rows`.
- Key: "tag50"
- Value: [Row 1, Row 101, Row 201...]

## The Trade-off
- **Read:** Extremely fast for "Contains" queries (`@>`).
- **Write:** Extremely slow. Updating a GIN index is expensive because one row change might require updating many keys in the index.

## Senior Tip
Only use JSONB/GIN if you truly need flexible schema. If your data structure is fixed, a normal table with columns is always faster and cleaner.
