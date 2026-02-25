# Week 7, Lesson 1: Share vs Exclusive Locks

## The Analogy: The Public Library

### 1. Share Lock (`FOR SHARE`)
- **Action:** Reading a book at the library table.
- **Rule:** Many people can read the same book at once.
- **Restriction:** No one can change the book (UPDATE) or throw it away (DELETE) while people are reading it.
- **Compatibility:** Share locks are "friends" with other Share locks.

### 2. Exclusive Lock (`FOR UPDATE`)
- **Action:** Taking the book to a private room to edit the pages.
- **Rule:** Only one person can have the book.
- **Restriction:** No one else can read it or edit it.
- **Compatibility:** Exclusive locks are "lonely." They don't share with anyone.

---

## When to use what?

| Scenario | Use `FOR SHARE` | Use `FOR UPDATE` |
| :--- | :--- | :--- |
| **Generating a Report** | Yes (Ensure data doesn't change) | No (Too slow, blocks others) |
| **Sending an Email with Data** | Yes | No |
| **Adding Money to Account** | No | Yes (Must be exclusive) |
| **Deleting a User** | No | Yes |

## The Conflict
If Terminal A has a `SHARE` lock and Terminal B has a `SHARE` lock, they are happy.
If Terminal C tries to `UPDATE`, it must wait until **BOTH** A and B are finished (`COMMIT` or `ROLLBACK`).
