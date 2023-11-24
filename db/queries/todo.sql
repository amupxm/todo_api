-- name: GetTodoByID :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTodo :one
INSERT INTO todos (task , parent_id)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET task = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

-- name: ToggleTodo :one
UPDATE todos
SET completed = NOT completed
WHERE id = $1
RETURNING *;
