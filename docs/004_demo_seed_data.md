# Seeding Demo Data

i'll need rows in my database to demonstrate how application will work.

> - Monthly Unique Users

largest granularity is monthly - so we should have at 1 month's worth of seed data

about 50 users a day would be the upper limit before generating data got too tedious

each user would login 0 - 3 times a day so we can demonstrate duplicate logins.

## Conclusion

- go seed script has the benefits (atomicity, easy to modify, deterministic) with
none of the downside (one large file that's hard to parse, might hit gh's limits)
