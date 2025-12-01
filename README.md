# Ticketing System (Go Fiber + MySQL)

Sistem ticketing sederhana dengan role **User** dan **Admin**. Dibangun
menggunakan **Golang (Fiber)** sebagai backend dan **MySQL** sebagai
database. Mendukung autentikasi JWT, role-based middleware, dan
manajemen tiket.

## Tech Stack

-   Backend: Go Fiber
-   Database: MySQL
-   ORM: GORM
-   Auth: JWT
-   Auto Reload: Air
-   Frontend: React.js (opsional)

## Project Structure

    ticketing-backend/
    │
    ├── main.go
    ├── go.mod
    ├── go.sum
    ├── .env
    ├── .env.example
    │
    ├── config/
    │   └── database.go
    │
    ├── controllers/
    │   ├── auth_controller.go
    │   ├── user_ticket_controller.go
    │   └── admin_ticket_controller.go
    │
    ├── middleware/
    │   └── jwt.go
    │
    ├── models/
    │   ├── user.go
    │   ├── ticket.go
    │   └── comment.go
    │
    └── routes/
        └── routes.go

## Installation

### 1. Clone Repository

``` sh
git clone https://github.com/yourusername/ticketing-backend.git
cd ticketing-backend
```

### 2. Install Dependencies

``` sh
go mod tidy
```

### 3. Install Air

``` sh
go install github.com/air-verse/air@latest
```

## Setup Environment

Buat file `.env`:

    APP_PORT=8080

    DB_USER=root
    DB_PASSWORD=your_password
    DB_HOST=localhost
    DB_NAME=pweb_ujian

    JWT_SECRET=supersecret123

## Run Application

### Dengan Air

``` sh
air
```

### Tanpa Air

``` sh
go run main.go
```

## Authentication & Roles

  Role    Akses
  ------- --------------------------------------------------------------
  User    Membuat tiket, melihat tiket milik sendiri, memberi komentar
  Admin   Melihat semua tiket, update status tiket

## API Routes

### Public

  Method   Route           Keterangan
  -------- --------------- -------------
  POST     /api/register   Daftar user
  POST     /api/login      Login →
