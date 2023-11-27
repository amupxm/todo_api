-- name: CreateUser :one
INSERT INTO users (username, password, email, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;
-- name: DeleteUser :one
UPDATE users
SET deleted_at = now()
WHERE id = $1
RETURNING *;
-- name: GetTodosByUser :many
SELECT *
FROM todos
WHERE user_id = $1 AND deleted_at IS NULL;
-- name: UpdateUser :one
UPDATE users
SET username = $2,
    password = $3,
    email = $4,
    first_name = $5,
    last_name = $6
WHERE id = $1
RETURNING *;
-- name: GetTodoByID :one
SELECT *
FROM todos
WHERE id = $1
AND  deleted_at IS NULL
LIMIT 1;
-- name: ListTodos :many
SELECT *
FROM todos
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: CreateTodo :one
INSERT INTO todos (user_id, title, task)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateTodo :one
UPDATE todos
SET title = $2,
    task = $3
WHERE id = $1
RETURNING *;
-- name: DeleteTodo :exec
UPDATE todos
SET deleted_at = now()
WHERE id = $1;
-- name: ToggleTodo :one
UPDATE todos
SET completed = NOT completed
WHERE id = $1
RETURNING *;