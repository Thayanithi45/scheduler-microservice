# Scheduler Microservice with GORM

A Golang microservice that allows job scheduling using GORM and PostgreSQL.

## Features
- Schedule recurring jobs
- REST API for job CRUD
- PostgreSQL via GORM
- Scheduler using robfig/cron

## Setup

```bash
go mod tidy
go run main.go
```

## API Endpoints
- `GET /jobs` - List jobs
- `GET /jobs/:id` - Get job by ID
- `POST /jobs` - Create job


