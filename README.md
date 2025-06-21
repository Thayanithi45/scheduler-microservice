# 🗓️ Scheduler Microservice

This project is a simple, production-style Golang microservice for managing scheduled jobs. It uses Gin for HTTP routing and GORM for database operations. Supports Docker + PostgreSQL integration.

## 📁 Project Structure

```
scheduler-microservice/
├── controllers/
├── database/
├── models/
├── docs/
├── main.go
├── go.mod / go.sum
├── .env
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## ⚙️ Tech Stack

- **Golang**
- **Gin** (HTTP framework)
- **GORM** (ORM)
- **PostgreSQL**
- **Docker & Docker Compose**
- **Swagger API Docs**

## 🔧 Environment Setup

Create a `.env` file in the project root:

```
PORT=8080
DATABASE_URL=postgres://postgres:postgres@db:5432/demodb?sslmode=disable
```

This is used by the app to connect to the database.

## 🚀 Run with Docker (Recommended)

> 🐳 No local Go/Postgres installation needed.

### 1️⃣ Build & Start

```
docker-compose up --build
```

### 2️⃣ App will start at:

- API: [http://localhost:8080](http://localhost:8080)
- Swagger: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## 🛠️ Run Without Docker (Local Go)

> Make sure PostgreSQL is running locally on port `5432`.

### 1️⃣ Update `.env`

```
PORT=8080
DATABASE_URL=postgres://postgres:postgres@localhost:5432/demodb?sslmode=disable
```

> Replace credentials and host as per your system.

### 2️⃣ Run with Go

```
go run main.go
```

## 📦 API Endpoints

| Method | Endpoint              | Description         |
|--------|-----------------------|---------------------|
| GET    | `/jobs/`              | Get all jobs        |
| GET    | `/jobs/detail/:id`    | Get job by ID       |
| POST   | `/jobs/create/`       | Create new job      |
| GET    | `/swagger/index.html` | Swagger API docs    |

## ✅ DB Credentials (Docker Setup)

These are set in `docker-compose.yml` and must match `.env`:

```yaml
environment:
  POSTGRES_USER: postgres
  POSTGRES_PASSWORD: postgres
  POSTGRES_DB: demodb
```

## ❓ FAQ

**Q: Where do I put the DB credentials?**  
👉 In the `.env` file.

**Q: How does Docker know them?**  
👉 `docker-compose.yml` configures the `db` container.

**Q: Why does Go fail with `lookup db: no such host`?**  
👉 You’re using `db` as host in `DATABASE_URL`, which only works *inside Docker*. Use `localhost` for local runs.

