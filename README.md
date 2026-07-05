# 🚗 SportSync Parking Reservation API

A RESTful Parking Reservation Management System built with Go, Echo, GORM, and PostgreSQL. The application provides secure JWT authentication, role-based authorization, parking zone management, and reservation management with concurrency-safe booking.

## 🌐 Live Demo

**API Base URL**

https://sportsync-server-go.onrender.com/

---

# ✨ Features

- JWT Authentication
- Role-Based Authorization (Admin & Driver)
- User Registration & Login
- Parking Zone Management
- Reserve Parking Spot
- Concurrency-safe reservation using Database Transactions and Row-Level Locking (`FOR UPDATE`)
- Cancel Reservation
- View My Reservations
- Admin Reservation Management
- Dynamic Available Spots Calculation
- Request Validation
- Layered Architecture (Handler → Service → Repository)

---

# 🛠 Tech Stack

- Go
- Echo v5
- GORM
- PostgreSQL
- JWT
- Validator
- Render (Deployment)

---

# 🏗 Project Architecture

The project follows a clean layered architecture.

```
                HTTP Request
                     │
                     ▼
                Echo Router
                     │
                     ▼
              Authentication Middleware
                     │
                     ▼
              Authorization Middleware
                     │
                     ▼
                  Handler Layer
                     │
                     ▼
                  Service Layer
          (Business Logic & Validation)
                     │
                     ▼
                Repository Layer
               (Database Operations)
                     │
                     ▼
                 PostgreSQL Database
```

### Layer Responsibilities

### Handler

- Parse request
- Validate request
- Read JWT claims
- Return HTTP response

### Service

- Business logic
- Permission checking
- Reservation rules
- Mapping Models ↔ DTOs

### Repository

- Database queries
- Transactions
- Row locking
- CRUD operations

---

# 📁 Project Structure

```
cmd/
    server/

internal/
    auth/
    config/
    database/
    middlewares/
    httpresponse/

    user/
        dto/

    parkingzones/
        dto/

    reservations/
        dto/

    server/

```

---

# ⚙️ Environment Variables

Create a `.env` file.

```env
DB_URL=postgres://username:password@localhost:5432/sportsync
JWT_SECRET=your_secret_key
PORT=5000
```

---

# 🚀 Run Locally

### Clone Repository

```bash
git clone https://github.com/mdsamiulislam54/SportSync-Server-GO

cd SportSync-Server-GO
```

### Install Dependencies

```bash
go mod tidy
```

### Run Project

```bash
go run ./cmd/server
```

or

```bash
air
```

---

# 📌 API Endpoints

## Authentication

| Method | Endpoint | Access |
|---------|----------|--------|
| POST | `/api/v1/auth/register` | Public |
| POST | `/api/v1/auth/login` | Public |

---

## Parking Zones

| Method | Endpoint | Access |
|---------|----------|--------|
| POST | `/api/v1/zones` | Admin |
| GET | `/api/v1/zones` | Authenticated |
| GET | `/api/v1/zones/:id` | Authenticated |
| PUT | `/api/v1/zones/:id` | Admin |
| DELETE | `/api/v1/zones/:id` | Admin |

---

## Reservations

| Method | Endpoint | Access |
|---------|----------|--------|
| POST | `/api/v1/reservations` | Driver, Admin |
| GET | `/api/v1/reservations/my-reservations` | Authenticated |
| DELETE | `/api/v1/reservations/:id` | Driver (Own), Admin |
| GET | `/api/v1/reservations` | Admin |

---

# 🔒 Authentication

Include JWT Token in the Authorization header.

```
Authorization: Bearer <your_token>
```

---

# 🚦 Role Permissions

| Role | Permissions |
|------|-------------|
| Admin | Manage zones, reserve parking, cancel any reservation, view all reservations |
| Driver | Reserve parking, cancel own reservation, view own reservations |

---

# 🔄 Reservation Concurrency

To prevent overbooking, reservations are processed using:

- Database Transaction
- Row-Level Locking (`FOR UPDATE`)
- Active Reservation Count Validation

This guarantees that when only one parking spot is available, only one request succeeds.

---

# 📦 Database

- PostgreSQL
- AutoMigrate using GORM
- Foreign Key Relationships
- Soft Delete (`gorm.Model`)

---

# 👨‍💻 Author

**Samiul Islam**

GitHub: https://github.com/mdsamiulislam54