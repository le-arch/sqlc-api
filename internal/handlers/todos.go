package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "todo-app/internal/db/sqlc"
)

func RegisterTodoRoutes(r chi.Router, q *db.Queries) {
	r.Post("/todos", createTodo(q))
	r.Get("/todos", listTodos(q))
	r.Get("/todos/{id}", getTodo(q))
	r.Patch("/todos/{id}", updateTodo(q))
	r.Delete("/todos/{id}", deleteTodo(q))
}

func createTodo(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Title string `json:"title"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		todo, _ := q.CreateTodo(r.Context(), db.CreateTodoParams{
			Title: body.Title,
			Description: sql.NullString{},
			Priority: 3,
		})
		json.NewEncoder(w).Encode(todo)
	}
}

func listTodos(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, _ := q.ListTodos(r.Context())
		json.NewEncoder(w).Encode(todos)
	}
}

func getTodo(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsedID, _ := uuid.Parse(id)
		todo, _ := q.GetTodo(r.Context(), parsedID)
		json.NewEncoder(w).Encode(todo)
	}
}

func updateTodo(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsedID, _ := uuid.Parse(id)
		var body struct {
			Completed bool `json:"completed"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		existing, _ := q.GetTodo(r.Context(), parsedID)
		updated, _ := q.UpdateTodo(r.Context(), db.UpdateTodoParams{
			ID: parsedID,
			Title: existing.Title,
			Description: existing.Description,
			Completed: body.Completed,
			Priority: existing.Priority,
		})
		json.NewEncoder(w).Encode(updated)
	}
}

func deleteTodo(q *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsedID, _ := uuid.Parse(id)
		q.DeleteTodo(r.Context(), parsedID)
		w.WriteHeader(http.StatusNoContent)
	}
}
