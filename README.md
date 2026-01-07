# Go Boilerplate Code

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)
![Gin Framework](https://img.shields.io/badge/Gin-v1.10-00ADD8?style=flat&logo=go)
![GORM](https://img.shields.io/badge/GORM-v1.25-2C5BB4?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)
![Redis](https://img.shields.io/badge/Redis-Latest-DC382D?style=flat&logo=redis)
## 🚀 Overview

GoBoilerPlate is a reusable backend foundation for building RESTful APIs and microservices in Go.

I built this project to avoid repeatedly re-implementing the same backend concerns—authentication, configuration, database access, caching, logging, and infrastructure wiring—across multiple projects.

The focus is not on features, but on **structure, boundaries, and maintainability**, following Clean Architecture principles to keep business logic independent of frameworks and infrastructure.

## 🎯 What This Project Demonstrates

- Clean Architecture and dependency boundaries in Go
- Modular API design with Gin
- JWT-based authentication flow
- Database access patterns using GORM
- Redis caching integration
- Environment-based configuration management
- Structured logging suitable for production services
- Testable service and repository layers

## ✨ Key Features

- **🏗 Clean Architecture**: Separation of concerns using `internal` (domain logic), `apis` (controllers), and `shared` (infrastructure) layers.
- **🔌 RESTful API with Gin**: High-performance HTTP web framework.
- **🗄 Database Management**:
  - **GORM**: Powerful ORM for PostgreSQL.
  - **Migrations**: CLI-based database version control using `golang-migrate`.
- **🔐 Authentication & Security**:
  - **JWT Authentication**: Secure stateless authentication strategy.
  - **Middleware**: Custom middlewares for auth, logging, and error handling.
- **⚡ Caching**: Integrated **Redis** client for high-speed caching.
- **☁️ File Storage**: AWS S3 integration for robust file handling.
- **📝 Configuration**: Environment-based config management using **Viper**.
- **📃 Logging**: Structured, high-performance logging with **Zap**.
- **🛠 Developer Experience**:
  - **Hot Reload**: Pre-configured with **Air** for live reloading during development.
  - **Makefile**: Handy commands for common tasks (run, build, test, migrate).

## 🛠 Tech Stack

- **Language**: [Go (Golang)](https://go.dev/)
- **Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM**: [GORM](https://gorm.io/)
- **Cache**: [Redis](https://redis.io/)
- **Config**: [Viper](https://github.com/spf13/viper)
- **Logging**: [Zap](https://github.com/uber-go/zap)
- **Validation**: [Go Playground Validator](https://github.com/go-playground/validator)
- **AWS SDK**: [AWS SDK for Go v2](https://aws.github.io/aws-sdk-go-v2/)

## 📂 Project Structure

```bash
.
├── apis                # API Layer (Routes, Controllers, DTOs)
│   ├── routes          # Route definitions
│   └── ...
├── cmd                 # Entry points
│   └── server          # Main server application
├── config              # Configuration files (config.yml)
├── internal            # core Application Logic (Domain, Service, Repository)
│   ├── auth            # Authentication module
│   ├── user            # User management module
│   ├── file            # File handling module
│   └── ...
├── shared              # Shared libraries (DB, Logger, Utils)
│   └── clients         # External clients (Database, Redis, AWS)
├── Makefile            # Task runner commands
└── go.mod              # Go module definitions
```

## ⚡ Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go 1.22+](https://go.dev/dl/)
- [Make](https://www.gnu.org/software/make/) (optional, but recommended)

### 1. Clone the Repository

```bash
git clone https://github.com/Prabhat070saini/coreGoboilerplatecode.git
cd coreGoboilerplatecode
```

### 2. Configuration

Copy the example configuration file and update it with your credentials:

```bash
cp config/example.config.yml config/config.yml
```

> **Note:** Update `config/config.yml` with your database (PostgreSQL), Redis, and AWS credentials.


### 3. Run Locally

If you prefer running Go locally:

**Step 3a: Start Dependencies**
Ensure PostgreSQL and Redis are running (e.g., via Docker):

```bash
docker run -d --name postgres-testing -e POSTGRES_DB=Testing -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres:16

docker run -d --name redis-cache -p 6379:6379 redis:latest

```

**Step 3b: Run Migrations**
Apply database migrations to set up the schema:

```bash
make migrate
```

**Step 3c: Start the Server**

```bash
make run
```

_Or for hot-reload during development:_

```bash
make run-watch
```

The server will start at `http://localhost:8080` (or your configured port).

## 🧰 Makefile Commands

We use `make` to simplify common development tasks:

| Command                     | Description                             |
| :-------------------------- | :-------------------------------------- |
| `make run`                  | Run the application normally            |
| `make run-watch`            | Run with hot-reloading (requires `air`) |
| `make migrate`              | Apply all up database migrations        |
| `make migrate-down`         | Rollback the last migration             |
| `make migrate-new name=...` | Create a new migration file             |
| `make migrate-reset`        | Drop all tables and re-run migrations   |
| `make test`                 | Run all tests                           |

## 🧪 Testing

Run almost all unit and integration tests:

```bash
make test
```

### Mocks Generation

We use `gomock` for mocking interfaces. To generate mocks:

```bash
mockgen -source=internal/user/service.go -destination=internal/user/mock/service_mock.go -package=mock
```

## 🤝 Contributing

Contributions are welcome! Please fork the repository and submit a Pull Request.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

---

_Built with ❤️ by [Prabhat Saini](https://github.com/Prabhat070saini)_
