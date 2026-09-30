<p align="center">
  <a href="https://nestjs.com/" target="_blank">
    <img src="https://nestjs.com/img/logo-small.svg" width="120" alt="NestJS Logo" />
  </a>
</p>

# NestJS Boilerplate Code

A reusable, production-oriented boilerplate for building scalable server-side applications using **NestJS** and **TypeScript**. Designed to standardize backend architecture and eliminate repeated setup of authentication, infrastructure, and cross-cutting concerns across projects.

---

## 🚀 Overview

This project was created to avoid repeatedly re-implementing the same backend foundations—authentication, configuration, database access, caching, queues, validation, and logging—across multiple services.

The focus is on **structure, maintainability, and real-world backend concerns**, not demo features. It serves as a solid starting point for enterprise-grade NestJS applications.

---

## 🎯 What This Project Demonstrates

- Modular NestJS architecture with clear separation of concerns
- Authentication and authorization using JWT and Passport
- Role-based access control (RBAC)
- Database access patterns with TypeORM and PostgreSQL
- Redis-based caching for performance optimization
- RabbitMQ-driven asynchronous processing
- Global validation, error handling, and logging
- Scalable, testable backend service design

---

## ✨ Key Features

- JWT-based authentication using Passport
- Google OAuth2 integration
- Role-based access control (RBAC)
- User management module with guarded routes
- PostgreSQL integration using TypeORM
- Database migrations and seeding support
- Redis integration for caching
- RabbitMQ for background jobs and async workflows
- File uploads using Multer
- Excel processing using `@univerjs-pro/exchange-client`
- Global validation pipes (`class-validator`, `class-transformer`)
- Centralized exception handling
- Custom logger for structured application logs

---

## 🛠 Tech Stack

- Framework: NestJS
- Language: TypeScript
- Database: PostgreSQL
- ORM: TypeORM
- Cache: Redis
- Message Queue: RabbitMQ

---
## 📂 Project Structure

- `src/apis`: Contains feature modules (Auth, User, Excel, etc.).
- `src/common`: Shared utilities, decorators, guards, and interceptors.
- `src/config`: Configuration files using `dotenv`.
- `src/shared`: Shared modules like Database, Queue, and Cache setup.

```text
src
├── apis                 # Feature specific modules
│   ├── auth             # Authentication module
│   ├── excel            # Excel processing module
│   └── user             # User management module
├── common               # Shared utilities
│   ├── constants        # Global constants
│   ├── decorators       # Custom decorators
│   ├── dto              # Shared DTOs
│   ├── guards           # Authentication guards
│   ├── lib              # Libraries (Logger, etc.)
│   ├── repository       # Base repositories
│   └── utils            # Helper functions
├── config               # Application configuration
└── Core               # Core  modules
    ├── cache            # Redis cache setup
    ├── database         # Database and TypeORM setup
    └── queue            # RabbitMQ queue setup
```
## ⚙️ Prerequisites

Ensure the following are installed and running locally:

- Node.js (v16 or higher)
- npm or yarn
- PostgreSQL
- Redis
- RabbitMQ

---

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/Prabhat070saini/CodeNestBoilerPlateCode.git
cd CodeNestBoilerPlateCode
```

```bash
npm install
```

```bash
npm run prepare
```
## ADD ENV (example env)
```
# ==========================
# App Config
# ==========================
X_API_KEY=your-api-key
ENV_NAME=development
LOG_LEVEL=debug
SERVER_HOST=localhost
SERVER_PORT=3000

# ==========================
# Database Config
# ==========================
DB_DIALECT=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=example_db
DB_POOL_MIN=5
DB_POOL_MAX=20
DB_CONNECTION_TIMEOUT=10000
DB_LOGGING=false
DB_CONNECTION_IDLE_TIMEOUT=60000

# ==========================
# Redis Config
# ==========================
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_CONNECTION_LIMIT=50
USE_REDIS=true
# REDIS_PASSWORD=

# ==========================
# Queue / RabbitMQ Config
# ==========================
RABBITMQ_URL=amqp://localhost
DEFAULT_RETRIES=3
DEFAULT_RETRY_DELAY_MS=5000
IS_DLQ=true

# ==========================
# JWT / Tokens
# ==========================
ACCESS_TOKEN_SECRET=your-access-token-secret
REFRESH_TOKEN_SECRET=your-refresh-token-secret
ACCESS_TOKEN_EXP_IN_MIN=15m
REFRESH_TOKEN_EXP_IN_MIN=7d

# ==========================
# AWS / S3 (Optional)
# ==========================
AWS_ACCESS_KEY_ID=your-aws-access-key
AWS_SECRET_ACCESS_KEY=your-aws-secret-key
AWS_REGION=us-east-1
AWS_S3_BUCKET_NAME=your-bucket-name
SIGNED_URL_EXPIRE=3600
# STORAGE_CDN_URL=https://cdn.example.com

# ==========================
# Email Config
# ==========================
EMAIL_HOST=smtp.example.com
EMAIL_PORT=465
EMAIL_USERNAME=example@example.com
EMAIL_PASSWORD=your-email-password
EMAIL_FROM=example@example.com
EMAIL_PROVIDER=smtp

# ==========================
# Google OAuth Config
# ==========================
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret

# ==========================
# OTP Config
# ==========================
OTP_CRYPTO_SECRET=your-otp-secret
OTP_TTL=300
COOL_DOWN_TTL=60
MAX_PER_HOUR=3
MAX_ATTEMPTS=3
```
## ▶️ Running the Application
```bash
npm run start
```
_Or for hot-reload during development:_

```bash
npm  run start:dev
```

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
<div align="center">

**⭐ If you find this project helpful, please give it a star!**

Made with ❤️ by [Prabhat Saini](https://github.com/Prabhat070saini)

</div>

