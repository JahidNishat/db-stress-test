# Week 7, Lesson 4: `SKIP LOCKED` (Building Queues)

## The Problem with SQL Queues
Traditionally, using a database as a queue was slow.
- Worker A locks Row 1.
- Worker B tries to lock Row 1 -> Wait -> Block.

## The Solution: `SKIP LOCKED`
- Worker A locks Row 1.
- Worker B sees Row 1 is locked -> **Skips it** -> Locks Row 2.
- **Result:** Zero waiting. Infinite scalability (up to row count).

## The Query
```sql
SELECT * FROM jobs
WHERE status = 'pending'
ORDER BY id ASC
LIMIT 1
FOR UPDATE SKIP LOCKED;
```

## Why `LIMIT 1`?
If you don't limit, Worker A will try to lock **every pending job** in the table!
By using `LIMIT 1`, you take only what you can chew.

## Use Cases
- **Email Sending:** "Give me the next email to send."
- **Video Transcoding:** "Give me the next video to process."
- **Order Fulfillment:** "Give me the next order to pack."
