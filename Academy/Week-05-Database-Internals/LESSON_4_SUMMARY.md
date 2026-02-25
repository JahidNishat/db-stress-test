# Week 5, Lesson 4: The Contention Penalty (Summary)

In our experiment, we found:
- **Pessimistic Locking:** ~180ms
- **Optimistic Locking:** ~400ms

## Why?
This is called **Contention**. When 100 people try to change the *same* row at the *same* time:
1.  **Pessimistic** creates a "Queue." The database manages the line efficiently.
2.  **Optimistic** creates a "Collision." 99 people fail, then 98 fail, then 97 fail... This results in thousands of unnecessary database calls and CPU retries.

## The Senior Rule of Thumb:
- **Use Pessimistic Locking** when you expect many people to fight over the same resource (e.g., a "Buy Now" button for a limited-stock item).
- **Use Optimistic Locking** when you expect overlaps to be rare (e.g., two admins editing different parts of a user's profile at the same time).
