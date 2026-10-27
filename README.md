# Book Management System

A REST API for managing books. The service is written in Go, uses MySQL for
persistent storage, and provides endpoints to create, retrieve, update, and
delete book records.

## Features

- Create a book with its name, author, and publication.
- List all books or retrieve one book by ID.
- Update the non-empty fields supplied for a book.
- Delete a book. GORM's default soft-delete behavior is enabled through
  `gorm.Model`.
- Automatically migrate the books table when the application starts.

## Technologies

- **Go** — application language (the module specifies Go `1.27.3`).
- **MySQL** — relational database.
- **GORM** — ORM and database migrations.
- **Gorilla Mux** — HTTP router.
- **godotenv** — loads local environment variables from `.env`.
- **go-sql-driver/mysql** — MySQL driver used by GORM.

## Project directory

```text
BookManagementSystem/
├── cmd/
│   └── main.go                       # Application entry point; starts HTTP server
├── pkg/
│   ├── config/
│   │   └── app.go                    # MySQL connection and environment configuration
│   ├── controllers/
│   │   └── book-controller.go        # HTTP handlers and request/response processing
│   ├── models/
│   │   └── book.go                   # Book model and database operations
│   ├── routes/
│   │   └── bookstore-routes.go       # Book API route registration
│   └── utils/
│       └── utils.go                  # Shared request parsing utility
├── go.mod                            # Module name and dependency versions
├── go.sum                            # Dependency checksums
└── README.md
```

## Prerequisites

- Go `1.27.3` or a compatible later version.
- A running MySQL server and a database created for this application.

## Configuration

Set the following environment variables before starting the service. They can
be provided by the shell or in a local `.env` file in the project root.

| Variable | Description |
| --- | --- |
| `DB_USER` | MySQL username |
| `DB_PASSWORD` | MySQL password |
| `DB_HOST` | MySQL server hostname |
| `DB_PORT` | MySQL server port, commonly `3306` |
| `DB_NAME` | Name of an existing MySQL database |

Example `.env` (replace the placeholder values; do not commit real credentials):

```dotenv
DB_USER=your_mysql_user
DB_PASSWORD=your_mysql_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=book_management
```

The application loads `.env` when present and otherwise uses the environment
already set by the operating system. The database itself must exist before
startup; GORM creates or updates the books table automatically.

## Run the application

From the project root:

```powershell
go run ./cmd
```

The server listens on `http://localhost:8000`.

## API

All endpoints use JSON request and response bodies where applicable.

| Method | Path | Description |
| --- | --- | --- |
| `POST` | `/book` | Create a book |
| `GET` | `/book` | List all books |
| `GET` | `/book/{bookId}` | Get a book by ID |
| `PUT` | `/book/{bookId}` | Update the non-empty fields in a book |
| `DELETE` | `/book/{bookId}` | Soft-delete a book |

### Book JSON

```json
{
  "name": "The Hobbit",
  "author": "J. R. R. Tolkien",
  "publication": "George Allen & Unwin"
}
```

`name`, `author`, and `publication` are required when creating a book. For
updates, provide the fields to change; empty or omitted values are left
unchanged. IDs must be positive integers.

### Example requests

Create a book:

```powershell
curl.exe -X POST http://localhost:8000/book `
  -H "Content-Type: application/json" `
  -d '{"name":"The Hobbit","author":"J. R. R. Tolkien","publication":"George Allen & Unwin"}'
```

List books or retrieve one book:

```powershell
curl.exe http://localhost:8000/book
curl.exe http://localhost:8000/book/1
```

Update a book:

```powershell
curl.exe -X PUT http://localhost:8000/book/1 `
  -H "Content-Type: application/json" `
  -d '{"publication":"Updated publisher"}'
```

Delete a book:

```powershell
curl.exe -X DELETE http://localhost:8000/book/1
```

Invalid IDs or JSON return `400 Bad Request`; a missing book returns
`404 Not Found`; unexpected database failures return `500 Internal Server
Error`. Successful create requests return `201 Created`; other successful
requests return `200 OK`.
