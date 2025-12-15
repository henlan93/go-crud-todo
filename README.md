# go-crud-todo

Simple CRUD API built with Go, Gin, and PostgreSQL(Supabase).

## Run

```bash
go mod tidy
go run main.go


.env:
DATABASE_URL=your_url
PORT=8080

Endpoints:
POST   /todos
GET    /todos/:id
PUT    /todos/:id
DELETE /todos/:id

