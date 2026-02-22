# Todo App Service (Go + MySQL)

A production-style REST API built in Go demonstrating clean architecture, MySQL integration, Dockerization, validation, and proper API response structuring.

This project follows backend engineering best practices including layered architecture, context propagation, graceful shutdown, and UTC-based time handling.

---

## Tech Stack

- Go 1.22
- MySQL 8
- Docker & Docker Compose
- Clean Architecture (Handler → Service → Repository)
- UUID (v4)
- Context-aware DB operations

---

## Project Structure

```
todo-app/
│
├── main.go          # Application entry point
├── internal/
│   ├── config/          # Environment configuration
│   ├── handler/         # HTTP handlers
│   ├── service/         # Business logic layer
│   ├── repository/      # Database abstraction
│   ├── model/           # Domain models
│   └── response/        # Standard API response format
│
├── migrations/          # SQL schema
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## Features

-  CRUD operations for Todos
-  MySQL integration
-  Clean layered architecture
-  Input validation (task cannot be empty)
-  Standardized API response format
-  Context-based DB operations
-  Graceful shutdown support
-  Dockerized setup
-  UTC-based time handling
-  Sorted task listing (by due date ASC)
-  Optional query filter for completed tasks

---

## API Response Format

All responses follow a consistent envelope structure.

### Success

```json
{
  "success": true,
  "data": { ... }
}
```

### Error

```json
{
  "success": false,
  "error": {
    "code": "INVALID_INPUT",
    "message": "task cannot be empty"
  }
}
```

---

## Setup & Run

### Using Docker (Recommended)

```bash
docker-compose up --build
```

API runs at:

```
http://localhost:8080
```

---

### Run Locally (Without Docker)

Make sure MySQL is running and environment variables are set:

```bash
export DB_HOST=localhost
export DB_USER=root
export DB_PASSWORD=root
export DB_NAME=todo
```

Then run:

```bash
go run cmd/server/main.go
```

---

## API Endpoints

### Create Todo

```
POST /todos
```

Request Body:

```json
{
  "task": "Finish assessment",
  "due_date": "2026-03-01T10:00:00Z"
}
```

---

### Get Todo

```
GET /todos/{id}
```

---

### List Todos

```
GET /todos
GET /todos?include_completed=true
```

Default behavior:
- Excludes completed tasks
- Sorted by due_date ASC

---

### Update Todo

```
PUT /todos/{id}
```

---

### Delete Todo

```
DELETE /todos/{id}
```

---

## Time Handling

- All timestamps are stored in UTC.
- API returns timestamps in ISO 8601 format.
- Timezone conversion should be handled by the client.

---

## Author

Nihal  
Backend Engineer | Go | Distributed Systems | Microservices

