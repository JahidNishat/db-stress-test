# Week 5: Database Internals - ACID Deep Dive

## What is ACID?

1.  **Atomicity:** "All or Nothing." If one part of a transaction fails, the whole thing is rolled back.
2.  **Consistency:** The database moves from one valid state to another.
3.  **Isolation:** Transactions don't interfere with each other.
4.  **Durability:** Once committed, data stays written even if the power fails.

---

## Lesson 1: The "Lost Update" Problem

This is a classic failure of **Isolation**. 

### The Scenario:
- Account Balance: $100.
- Transaction 1: Reads $100, plans to add $50.
- Transaction 2: Reads $100, plans to add $20.
- Transaction 1: Writes $150.
- Transaction 2: Writes $120.
- **Result:** The balance is $120. The $50 from Transaction 1 is **LOST**.

---

## Task: Visualizing the Lost Update

1.  We will use our `db-stress` tool (the one we built earlier) or a new script.
2.  We will launch 100 goroutines, each trying to increment a counter by 1.
3.  If we don't use proper transactions/locking, the final result will NOT be 100.

### Implementation Checklist
- [ ] Create `bank.sql` with an `accounts` table.
- [ ] Write `lost_update.go` to simulate the race condition.
- [ ] Run the test and see the failure.
