package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"

	db "todo-app/internal/db/sqlc"
	"todo-app/internal/handlers"
)

func main() {
	godotenv.Load()

	dbConn, _ := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	queries := db.New(dbConn)

	r := chi.NewRouter()
	handlers.RegisterTodoRoutes(r, queries)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
