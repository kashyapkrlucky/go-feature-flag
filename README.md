# Feature Flag Service

This repository implements a Feature Flag Service for managing feature toggles in a cloud environment. It allows teams to control the rollout of features dynamically using feature flags, enabling continuous delivery practices and gradual feature deployment.

---

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Technologies](#technologies)
- [System Design](#system-design)
- [Setup](#setup)
- [Development](#development)
- [Testing](#testing)
- [File Structure](#file-structure)
- [What Was Done](#what-was-done)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

Feature flags (also known as feature toggles) are a powerful tool for managing and controlling features in production. They allow you to toggle specific features on or off at runtime, without needing to redeploy the application.

This service provides APIs to manage feature flags, such as creating, updating, deleting, and checking the state of a feature flag. It also integrates Redis for caching feature flag data to improve performance and PostgreSQL for persistent storage.

---

## Features

- **Create, Read, Update, Delete (CRUD) operations** for feature flags.
- **Caching** with Redis for faster retrieval of flag states.
- **Integration with PostgreSQL** for persistent storage of feature flags.
- **API rate limiting** to control the number of requests.
- **Role-based access control (RBAC)** for managing permissions to create and manage flags.
- **REST API** to interact with the service, including endpoints for:
  - Create Feature Flag
  - Update Feature Flag
  - Get Feature Flag
  - Delete Feature Flag

---

## Technologies

- **Backend Framework**: Go (Golang)
- **Database**: PostgreSQL (for storing feature flags)
- **Cache**: Redis (for caching flag states)
- **API Framework**: RESTful APIs
- **Testing**: Unit and Integration tests with `GoTest`, `Testify`, `miniredis`
- **CI/CD**: GitHub Actions (or other tools like Jenkins)
- **Containerization**: Docker (for development and production environments)
- **Container Orchestration**: Kubernetes (for deployment)
- **API Documentation**: Swagger (optional)

---

## System Design

The system is designed to provide scalable, fast, and reliable management of feature flags. It follows a **microservices architecture** where the feature flag service runs independently and interacts with other services through well-defined REST APIs.

### Components:

1. **Feature Flag Service**:

   - Handles CRUD operations for feature flags.
   - Caches the results using Redis for quick retrieval.

2. **PostgreSQL Database**:

   - Stores feature flag data persistently.
   - Supports querying and updating flags.

3. **Redis Cache**:

   - Stores feature flag data temporarily to reduce latency.
   - Caching improves performance by storing the flag state.

4. **API Gateway** (optional):
   - Manages routing and serves as an entry point to the system.

---

## Setup

To get started with this project, follow these steps:

### 1. Clone the repository:

```bash
git clone https://github.com/kashyapkrlucky/feature-flag-service.git
cd feature-flag-service
```

### 2. Install dependencies:

Ensure you have Go installed on your local machine. You can install Go from [here](https://golang.org/dl/).

Install necessary Go packages:

```bash
go mod tidy
```

### 3. Set up PostgreSQL:

Create a PostgreSQL database (e.g., `feature_flags_db`) and configure your environment variables:

```bash
export DB_USER=your_db_user
export DB_PASSWORD=your_db_password
export DB_NAME=feature_flags_db
```

### 4. Set up Redis:

Install and run Redis locally or use a cloud-based Redis service. Update your Redis connection settings:

```bash
export REDIS_HOST=localhost
export REDIS_PORT=6379
```

### 5. Run the application:

Now, run the application locally:

```bash
go run main.go
```

---

## Development

To contribute or modify the service, follow these steps:

### 1. Fork the repository:

Fork the repository to your own GitHub account.

### 2. Create a new branch:

Create a new feature or bugfix branch:

```bash
git checkout -b feature/my-new-feature
```

### 3. Make changes:

Implement your changes, following the coding standards and best practices.

### 4. Commit changes:

Make sure to write clear and concise commit messages.

```bash
git commit -m "Add feature flag validation"
```

### 5. Push changes:

Push your changes to your fork:

```bash
git push origin feature/my-new-feature
```

### 6. Create a pull request:

Open a pull request from your feature branch to the main repository.

---

## Testing

To run the tests, simply use the `go test` command:

```bash
go test ./internal/repositories -v
```

You can also run all tests across the repository:

```bash
go test ./...
```

---

## File Structure

Here’s a breakdown of the key files and directories in the project:

### `main.go`

- The entry point of the application, where the server is initialized and started.

### `internal/repositories/`

- Contains the data access layer for interacting with the PostgreSQL database and Redis.
  - `feature_flag_repo.go`: Contains methods to interact with the `feature_flags` table in PostgreSQL.
  - `feature_flag_repo_test.go`: Unit and integration tests for the `feature_flag_repo`.

### `internal/models/`

- Defines the data models used by the application, including `FeatureFlag`.
  - `feature_flag.go`: The data structure for a feature flag.

### `db/`

- Contains database-related utility functions.
  - `db.go`: Initializes the database connection to PostgreSQL.
  - `redis.go`: Initializes the Redis connection.

### `api/`

- Contains the HTTP handlers for the REST API endpoints related to feature flag operations.
  - `feature_flag_handler.go`: Implements HTTP handlers for CRUD operations on feature flags.

### `go.mod`

- The Go module definition file, which lists dependencies for the project.

---

## What Was Done in the Code

Here’s an overview of the tasks completed and what was done in the code:

### 1. **PostgreSQL Integration**:

- A PostgreSQL database is used to store feature flag data persistently.
- Created a `feature_flags` table with columns like `id`, `name`, and `enabled`.

### 2. **Redis Caching**:

- Integrated Redis to cache the results of feature flag queries for faster retrieval.
- Used `miniredis` for mocking Redis in unit tests.

### 3. **Feature Flag CRUD Operations**:

- Implemented the core feature flag operations, including create, update, delete, and get.
- Created methods for interacting with both Redis and PostgreSQL.

### 4. **Unit and Integration Tests**:

- Wrote unit tests for repository methods using `Testify` for assertions.
- Used `miniredis` to mock Redis during tests to avoid hitting the real Redis server.

### 5. **Testing Environment Setup**:

- Set up a mock PostgreSQL database and Redis server for testing purposes.

### 6. **REST API**:

- Developed RESTful API endpoints to interact with feature flags.
- Created a handler that manages the CRUD operations for feature flags.

---

## Contributing

We welcome contributions to improve the service. If you find a bug or have a feature suggestion, please open an issue or create a pull request. Make sure your code passes all tests and is properly documented.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
