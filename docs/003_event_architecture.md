# Events

in most event-driven systems, events would be streamed
(e.g. rabbit, kafka, sqs, redis stream, etc) with a dead letter queue.

however, given them time constraint of 3 - 5 hours, that seems like
over-kill.

sticking with the PostgreSQL solution, here are some considerations:

- LISTEN/NOTIFY
  - lack of DLQ means this is a NO GO
- append-only table
  - seems to be what is being recommended by the requirements doc

in terms of performance, it seems in-efficient to compute each, GET request
by querying the entire table. especially since users cannot login in the past -
old data can be pre-computed and the results can be stored.

one possible solution could be to store old data (X amount time before now)
via redis, materialized view or otherwise so that we do not need to recompute.

the newer data (e.g. today / yesterday) can be computed by scanning the
user_logins table. i.e. decide which form of storage to query from based on
the age of data requested.

## Conclusion

i'm overthinking the problem. a suggestion was given in the requirements file,
and only 3 - 5 hours are given.

go with the simple solution:

- append-only table
- simple `SELECT COUNT(DISTINCT ...) ... WHERE date` queries should be fine
