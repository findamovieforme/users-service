# Users Service

A Go microservice responsible for user preferences and movie-related data in the Find a Movie for Me ecosystem.

## Design & Architecture

### Overview

The service follows a **layered architecture**: HTTP requests are handled by **routers**, delegated to **handlers** for request/response handling, then to **services** for business logic. Data is persisted in **AWS DynamoDB**.

```
HTTP → Router → Handler → Service → DynamoDB
```

### API Endpoints Exposed

- **Preferences** — Store and retrieve per-user movie preferences (`POST`/`GET /users/preferences`).
- **Most Added** — Aggregate across users to return the most frequently added movies (`GET /users/mostAdded`).
- **Health** — Root and `/users/ping` for liveness/readiness.

### Data & Infrastructure

- **Users**: User identity and authentication are handled by **AWS Cognito**. This service does not manage user accounts; it consumes user identifiers from AWS and stores preferences keyed by them.
- **Storage**: AWS DynamoDB, using a table named `UserPreferences` keyed by `user_id`, with a flexible `preferences` payload.
- **Runtime**: Configuration is driven by environment variables (e.g. `AWS_REGION`); optional `.env` is supported via `godotenv`.
