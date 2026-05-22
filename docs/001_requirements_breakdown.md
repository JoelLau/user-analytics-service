# Requirements Breakdown

## Unit Tests

> Include unit tests to validate logic and handle edge cases such as:

- tests are a **REQUIREMENT**.
- keep test coverage moderately high
- ~~consider using CI/CD to track coverage~~ (> 3 - 5 hours)
- consider using TDD
- we'll probably need "e2e" tests using testcontainers to check end-to-end flow
(use testcontainers to populate the db for each test)

## Cloud-Aware Architecture

> stateless services

i will not perform the count in memory on the Go server

> clear seperation of concerns

organize the code properly

> REST API structure

perfect for OpenAPI codegen

> containerization ... appreciated

don't include docker-compose in gitignore

## Database Design

> Ensure accurate uniqueness (no double-counting).

- sounds like a simple `UNIQUE` clause?

> Support performant aggregation queries for daily and monthly views.

- sounds like indexes are required

---

While the requirements doc seems to suggest use of relational
databases, this sounds like the wrong tool?

Depending on scale, I think we might need some sort of event queue
(e.g. kafka, etc) and workers to ingest user login events.

More research is required here.

## Expected Outputs

> Implement endpoints or functions that provide:

I may have been to eager to jump to the conclusion that a REST API is required.

> How to run the service

"the service" seems to be implied that its a long running process?

> Design decisions and assumptions.

write design and decision docs

## Conclusions

- write design docs
- create a http server, name the handler functions as recommended
- include tests where possible
