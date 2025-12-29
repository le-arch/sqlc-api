-- name: CreateTodo :one
INSERT INTO todos (title, description, priority)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos ORDER BY created_at DESC;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, description = $3, completed = $4, priority = $5, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;
