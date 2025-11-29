# Ticketing System (Go Fiber + MySQL)

Sistem ticketing sederhana dengan role **User** dan **Admin**. Dibangun menggunakan **Golang (Fiber)** sebagai backend dan **MySQL** sebagai database. Mendukung autentikasi JWT, role-based middleware, dan manajemen tiket.

---

## 🚀 Tech Stack

* **Backend**: Go Fiber
* **Database**: MySQL
* **ORM**: GORM
* **Auth**: JWT
* **Auto Reload**: Air
* **Frontend**: React.js (optional)

---

## 📁 Project Structure

```
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
```

---

## ⚙️ Installation

### 1. Clone Repo

```sh
git clone https://github.com/yourusername/ticketing-backend.git
cd ticketing-backend
```

### 2. Install Dependencies

```sh
go mod tidy
```

### 3. Install Air (auto reload)

```sh
go install github.com/air-verse/air@latest
```

---

## 🗂️ Setup Environment

Buat file `.env`:

```
APP_PORT=8080

DB_USER=root
DB_PASSWORD=your_password
DB_HOST=localhost
DB_NAME=pweb_ujian

JWT_SECRET=supersecret123
```

---

## ▶️ Run App (With Air)

```sh
air
```

Atau tanpa Air:

```sh
go run main.go
```

---

## 🔐 Authentication & Roles

Sistem mendukung 2 role:

| Role      | Akses                                                        |
| --------- | ------------------------------------------------------------ |
| **User**  | Membuat tiket, melihat tiket milik sendiri, memberi komentar |
| **Admin** | Melihat semua tiket, update status tiket                     |

---

## 📌 API Routes

### Public

| Method | Route           | Keterangan  |
| ------ | --------------- | ----------- |
| POST   | `/api/register` | Daftar user |
| POST   | `/api/login`    | Login → JWT |

---

### User

| Method | Route                           | Keterangan            |
| ------ | ------------------------------- | --------------------- |
| GET    | `/api/user/tickets`             | List tiket milik user |
| POST   | `/api/user/tickets`             | Membuat tiket         |
| GET    | `/api/user/tickets/:id`         | Detail tiket user     |
| POST   | `/api/user/tickets/:id/comment` | Menambah komentar     |

---

### Admin

| Method | Route                           | Keterangan          |
| ------ | ------------------------------- | ------------------- |
| GET    | `/api/admin/tickets`            | Semua tiket         |
| GET    | `/api/admin/tickets/:id`        | Detail tiket        |
| PUT    | `/api/admin/tickets/:id/status` | Update status tiket |

---

## 📘 Models

### User

```
ID
Name
Email
Password
Role
```

### Ticket

```
ID
Title
Message
Status
UserID
Comments[]
```

### Comment

```
ID
TicketID
UserID
Message
```

---

## 📝 Todo (Optional Features)

* Frontend React.js (User & Admin Dashboard)
* Pagination list ticket
* Notifikasi email
* Upload lampiran pada tiket
* Activity log

---

## 🤝 Author

**Muhamad Danendra Prawiraamijoyo**
