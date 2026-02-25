# Week 7, Lesson 3: Advisory Locks (Custom Locks)

## The Analogy: The Room Occupancy Sign

### 1. The Concept
An **Advisory Lock** is a lock that YOU define. Postgres doesn't know what it means. It's just a number (64-bit bigint).

### 2. Transaction vs Session
- **`pg_advisory_xact_lock(N)`**: Automatically unlocks when the transaction commits. (Safe).
- **`pg_advisory_lock(N)`**: Stays locked until you manually unlock it. (Dangerous, can leak locks).

---

## When to use what?

| Feature | `FOR UPDATE` | Advisory Lock |
| :--- | :--- | :--- |
| **Protects Data?** | Yes, physically locks the row. | No, only blocks others who *ask* for the same number. |
| **Requires Row?** | Yes, the row must exist. | No, you can lock any number. |
| **Performance** | Slower (updates indexes/mvcc). | Extremely fast (stored in memory). |

## Use Cases:
1. **Background Jobs:** Ensure only one worker is processing a specific "Job ID" or "Customer ID."
2. **Distributed Mutex:** Replace Redis locks if you already have a Postgres DB.
3. **Data Migrations:** Ensure only one instance of your migration script runs across a cluster.
