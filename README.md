# Service Catalog API 

This is a fault-tolerant, high-performance server written in Go using the **Gin framework**, **GORM** for ORM, and an **in-memory cache** for fast read access.

---

## Features

- `GET /services`
  - List services
  - Supports: `search`, `sort`, `limit`, `offset`

- `GET /services/:id`
  - Fetch single service by ID

- `GET /services/:id/versions`
  - Fetch versions of a service

---

## Design

```
Client ──> Gin Router ──> In-Memory Cache ──> SQLite DB (on startup)
```

- Data is seeded once and persisted in SQLite.
- On server start, data is loaded from DB into cache.
- All reads are served from the in-memory map with RWMutex protection.

---

## Tech Stack

- Go 1.24.5
- Gin Web Framework
- GORM ORM with SQLite
- `sync.RWMutex` for thread-safe caching
- `testify` for unit testing

---

## Project Structure

```
.
├── main.go                  # Entry point
├── go.mod                   # Module config
├── README.md                # Project overview
├── service_api_test.go      # Unit tests
└── internal/
    ├── db.go                # DB init and seed
    ├── cache.go             # Cache manager
    ├── handlers.go          # HTTP handlers
    └── models.go            # Data models
```

---

## Running Locally

```bash
go mod tidy
go run main.go
```

Visit:
- `http://localhost:8080/services`

---

## Testing

```bash
go test -v
```

Tests include:
- Search + Pagination
- Get service by ID
- Get versions by service
- Invalid and missing IDs

---

## Issues

| Issue                | Notes |
|---------------------|-------|
| Cache staleness     | No live DB sync; cache loaded once at startup |
| Cache invalidation  | No cache invalidation strategy |
| Cache size          | No cache size limit, can be fixed and need to add cache eviction policy (LRU, LFU, etc.)|

---

## Assumptions

- This is a **read-only** service with fixed data.
- Data consistency isn't a concern due to in-memory snapshot.

---

## Enhancements

- Add full CRUD APIs
- Periodic cache refresh / invalidation strategy and eviction policy
