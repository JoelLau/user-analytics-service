# Tech Stack

What external dependencies do I need?

## Considerations

- a **relational database** is always a good default place to start
  - i haven't figured out how event ingestion will happen,
  so i'll keep things generic / open-ended for now
  - postgresql vs mysql
    - problem description lists postgresql but job posting mentions mysql
      - i'll pick the one i'm more familiar with (postgresql)
- Go libraries
  - stdlib vs chi vs gin
    - wants: logging, middleware, recovery, param handling
    - i pick **chi** out of familiarity and because its handlers are compatible
      with stdlib (reversible decision)
  - REST APIs -> openapi documentation is best practice
    - make documentation source of truth by using openapi codegen
  - database access
    - sqlc vs ORMs vs raw drivers (pgx)
      - pgx is too unwieldy and hard to maintain if used raw
      - Go's ORMs are unlike Rails and a little weak,
      they don't provide enough value to justify "learning" a new wrapper
      (subjective decision)
      - SQLC allows optimized queries and produces type-safe code
        - will skip "repository layer" that wraps SQLC in the interest of time
- database set up:
  - credentials / roles are out of scope
  - migrations are part of requirements / documentation (also used by sqlc)
  - using Goose as migration tool so i can seed demo data
    - no "down" migrations because we're never going to revert

## Conclusions

- use **postgres** for persistence layer
- use **chi** as web server framework
- use **OpenAPI Codegen** to make documentation source of truth
- use **sqlc** to manage database access
