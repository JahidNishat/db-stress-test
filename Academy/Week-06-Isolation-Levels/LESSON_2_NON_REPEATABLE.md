# Week 6, Lesson 2: Non-Repeatable Reads

## The Phenomenon
A **Non-Repeatable Read** occurs when a transaction reads the same row twice and gets two different values because another transaction committed a change in between.

## Why it happens in "Read Committed"
In the default `Read Committed` mode, every single `SELECT` statement sees a new "snapshot" of the database. 
- Statement 1 sees the world at 12:00:00.
- Statement 2 sees the world at 12:00:01.

If a change was committed at 12:00:00.5, your two statements in the **same transaction** will disagree.

## The Experiment
1. **Terminal 1 (Reporter):** `BEGIN;` -> `SELECT balance...` (sees 100).
2. **Terminal 2 (Updater):** `UPDATE...` -> `COMMIT;` (sets to 150).
3. **Terminal 1 (Reporter):** `SELECT balance...` (sees 150!).

## The Danger
Imagine a "Total Assets" report:
1. Read Alice: $100.
2. (Alice transfers $50 to Bob).
3. Read Bob: $150 (instead of $100).
4. **Result:** Your report says the bank has $250 total, but it actually has $200. $50 was "created" by the report's inconsistency.
