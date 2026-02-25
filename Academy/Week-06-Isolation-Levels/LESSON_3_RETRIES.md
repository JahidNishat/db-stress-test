# Week 6, Lesson 3: Serialization Failures & Retries

## The Conflict
In `REPEATABLE READ` or `SERIALIZABLE` mode, Postgres ensures that your transaction is consistent. If you try to update a row that was modified and committed by another transaction **after** your transaction started, Postgres cannot guarantee consistency.

## The Error
You will see:
`ERROR: could not serialize access due to concurrent update`
(Postgres Error Code: `40001`)

## The Meaning
Postgres is saying: "I can't let you update this. The version of the row you are looking at is 'stale' (old). If I let you update it, you would be overwriting a change you never saw."

## The "Senior" Solution: Retries
Unlike `READ COMMITTED` where the database handles the wait, in `REPEATABLE READ` the **Application** must handle the failure.
1. Catch the error.
2. If the error code is `40001`, **Retry the entire transaction**.
