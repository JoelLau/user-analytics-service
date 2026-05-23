# Change to Recommended Schema

recommended schema from [REQUIREMENTS.md]

```sql
CREATE TABLE user_logins (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    login_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    UNIQUE(user_id, login_time)
);
```

current schema from [migrations/001_user_logins.sql]

```sql
CREATE TABLE user_logins (
    user_id      BIGINT      NOT NULL,
    logged_in_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

## removed `UNIQUE` constraint

i assume that in an events based system, another process is
responsible for generating _user login_ events.

the `user_logins` table should not lose any data and accept
what is generated.

the `UNIQUE` constraint prevents entries from being written
and results in an error on the event generator's part.
this should not be part of the responsibility of the caller.

## removed `id` column

there is nothing in any known system that requires a reference
to invididual rows in this table. they do not require their own
identity.

## _user_id_ `BIGINT`

the type was change for convenience in writing the seed script
and to make user_ids more human-readable for this demo.
