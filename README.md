# go-crud-todo

Simple CRUD API built with Go, Gin, and PostgreSQL(Supabase).

## Run

```bash
go mod tidy
go run main.go


.env:
DATABASE_URL=your_url
PORT=8080
API_TOKEN=your_token

Endpoints:
POST   /todos
GET    /todos/:id
PUT    /todos/:id
DELETE /todos/:id

