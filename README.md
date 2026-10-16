# Student API

A lightweight RESTful API for managing student records, built with **Go**. It uses the standard library for HTTP routing, **SQLite** for persistence, and request validation with a clean, layered project structure.

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-003B57?logo=sqlite&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [API Reference](#api-reference)
- [Development](#development)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- RESTful CRUD endpoints for student records
- Request payload validation with clear error responses
- Environment/YAML-based configuration
- Embedded SQLite database (pure Go driver, no CGO required)
- Clean architecture with separated `cmd`, `config`, and `internal` layers

## Tech Stack

| Layer          | Technology                                                              |
| -------------- | ----------------------------------------------------------------------- |
| Language       | Go                                                                      |
| Database       | SQLite via [`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) |
| Validation     | [`go-playground/validator`](https://github.com/go-playground/validator)  |
| Configuration  | [`cleanenv`](https://github.com/ilyakaznacheev/cleanenv)                 |

## Project Structure

```
Student-api/
├── cmd/
│   └── student-api/     # Application entry point
├── config/              # Configuration files
├── internal/            # Private application code (handlers, storage, types, etc.)
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.22 or newer
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/loveyadav8478-blip/Student-api.git
cd Student-api

# Download dependencies
go mod download
```

### Run the server

```bash
go run cmd/student-api/main.go -config config/local.yaml
```

The server starts on the address defined in your config (e.g. `http://localhost:8082`).

### Build a binary

```bash
go build -o bin/student-api ./cmd/student-api
./bin/student-api -config config/local.yaml
```

## Configuration

Configuration is loaded via `cleanenv` from a YAML file, and can be overridden with environment variables.

Example `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

| Key                   | Description                            | Example                 |
| --------------------- | -------------------------------------- | ----------------------- |
| `env`                 | Runtime environment                    | `dev`                   |
| `storage_path`        | Path to the SQLite database file       | `storage/storage.db`    |
| `http_server.address` | Host and port the server listens on    | `localhost:8082`        |

## API Reference

Base URL: `http://localhost:8082`

| Method | Endpoint             | Description              |
| ------ | -------------------- | ------------------------ |
| POST   | `/api/students`      | Create a new student     |
| GET    | `/api/students`      | List all students        |
| GET    | `/api/students/{id}` | Get a student by ID      |

### Create a student

**Request**

```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Love Yadav",
    "email": "love@example.com",
    "age": 20
  }'
```

**Response** `201 Created`

```json
{
  "id": 1
}
```

### Get a student

```bash
curl http://localhost:8082/api/students/1
```

```json
{
  "id": 1,
  "name": "Love Yadav",
  "email": "love@example.com",
  "age": 20
}
```

### Error format

```json
{
  "status": "Error",
  "error": "field name is a required field"
}
```

## Development

```bash
# Format code
go fmt ./...

# Static analysis
go vet ./...

# Run tests
go test ./...
```

## Roadmap

- [ ] Update and delete endpoints
- [ ] Pagination and filtering on list endpoint
- [ ] Dockerfile and docker-compose setup
- [ ] Unit and integration tests
- [ ] OpenAPI/Swagger documentation
- [ ] GitHub Actions CI pipeline

## Contributing

Contributions are welcome.

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Commit your changes: `git commit -m "Add your feature"`
4. Push to the branch: `git push origin feature/your-feature`
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for details.

## Author

**Love Yadav** · [GitHub](https://github.com/loveyadav8478-blip)