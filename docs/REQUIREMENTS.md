# Take-Home Assignment: User Analytics Service in Go

## Objective

Design and implement a backend service in **Go** that counts:

- Daily Unique Users
- Monthly Unique Users

You will simulate this using a persistent database. The emphasis is on clean design,
correctness, and production-readiness—especially for cloud-deployable services.

## Problem Statement

You're tasked with building a backend system that tracks user logins and provides:

- The count of **unique users** for each **day**.
- The count of **unique users** for each **month**.

Your system should support:

- Efficient ingestion of user activity logs.
- Queries for:
  - Daily unique users for a given day.
  - Monthly unique users for a given month.

## Requirements

### Implementation

- Use the **Go** programming language.
- Use a cloud-compatible relational database (e.g., PostgreSQL).
- Define and implement the data model and service logic.
- Include unit tests to validate logic and handle edge cases such as:
- Duplicate logins from the same user.
- Logins across day/month boundaries.
- Timezone handling if applicable.

### Cloud-Aware Architecture

Your solution should be structured with cloud deployment in mind (e.g., stateless
services, clear separation of concerns, RESTful API structure if applicable).
Containerization is optional but appreciated (e.g., Docker).

## Database Design

Provide a clear explanation of how you would structure the database to:

- Track user login events with timestamps.
- Ensure accurate uniqueness (no double-counting).
- Support performant aggregation queries for daily and monthly views.

Include a sample SQL schema. For example:

```sql
CREATE TABLE user_logins (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    login_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    UNIQUE(user_id, login_time)
);
```

You can refine this schema as you see fit. Justify your design and indexing strategy.

## Expected Outputs

Implement endpoints or functions that provide:

```plaintext
GetDailyUniqueUsers(date string) → int
GetMonthlyUniqueUsers(month string) → int
```

The actual interface design is flexible, but usage examples should be provided.

## Submission Checklist

- [ ] Go source code with modular structure.
- [ ] README.md explaining:
- [ ] How to run the service.
- [ ] Design decisions and assumptions.
- [ ] Sample input/output for queries.
- [ ] SQL schema and explanation.
- [ ] Unit tests and sample data.

## Time Estimate

Estimated effort: 3–5 hours
