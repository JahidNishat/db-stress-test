# Week 6 Project: Safe Transfer with Isolation Levels

## Goal
Build a Go script that performs a money transfer using `REPEATABLE READ` isolation and handles serialization conflicts with a retry loop.

## The Logic
1. Start a transaction.
2. `SET TRANSACTION ISOLATION LEVEL REPEATABLE READ`.
3. Read Alice's balance.
4. Read Bob's balance.
5. Subtract $50 from Alice, Add $50 to Bob.
6. Commit.

## The Conflict Test
To test it, run two instances of your script at the exact same time. One should succeed, and the other should print "Conflict detected, retrying..." and then eventually succeed.

## Key Learning
This is how modern high-consistency systems (like CockroachDB or Google Spanner) work. They prefer "Optimistic" concurrency at the protocol level and force the application to retry.
