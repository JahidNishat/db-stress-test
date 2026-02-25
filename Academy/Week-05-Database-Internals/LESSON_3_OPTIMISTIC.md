# Week 5, Lesson 3: Optimistic Locking

## 1. Pessimistic vs Optimistic
- **Pessimistic (`FOR UPDATE`):** Locks the row. Others must wait. Good for high-contention (Flash sales).
- **Optimistic (Version Check):** No locks. Checks if the version changed at the last second. Good for low-contention (User profile updates).

## 2. How to implement in SQL
```sql
UPDATE accounts 
SET balance = 150, version = version + 1 
WHERE id = 1 AND version = 1;
```
If `rowsAffected == 0`, someone else updated the row first. You must **Retry**.

## 3. Trade-offs
- **Pessimistic:** Guaranteed success but slower (queuing).
- **Optimistic:** Very fast but can fail under high load (needs retries).
