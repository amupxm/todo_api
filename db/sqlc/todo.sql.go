// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: todo.sql

package db

import (
	"context"
	"database/sql"
)

const CreateTodo = `-- name: CreateTodo :one
INSERT INTO todos (title,task)
VALUES ($1, $2)
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at
`

type CreateTodoParams struct {
	Title string
	Task  sql.NullString
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.queryRow(ctx, q.createTodoStmt, CreateTodo, arg.Title, arg.Task)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Task,
		&i.Completed,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const DeleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteTodoStmt, DeleteTodo, id)
	return err
}

const GetTodoByID = `-- name: GetTodoByID :one
SELECT id, title, task, completed, due_date, created_at, updated_at, deleted_at FROM todos
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodoByID(ctx context.Context, id int32) (Todo, error) {
	row := q.queryRow(ctx, q.getTodoByIDStmt, GetTodoByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Task,
		&i.Completed,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const ListTodos = `-- name: ListTodos :many
SELECT id, title, task, completed, due_date, created_at, updated_at, deleted_at FROM todos
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListTodosParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListTodos(ctx context.Context, arg ListTodosParams) ([]Todo, error) {
	rows, err := q.query(ctx, q.listTodosStmt, ListTodos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Task,
			&i.Completed,
			&i.DueDate,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ToggleTodo = `-- name: ToggleTodo :one
UPDATE todos
SET completed = NOT completed
WHERE id = $1
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at
`

func (q *Queries) ToggleTodo(ctx context.Context, id int32) (Todo, error) {
	row := q.queryRow(ctx, q.toggleTodoStmt, ToggleTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Task,
		&i.Completed,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const UpdateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $2
, task = $3
WHERE id = $1
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at
`

type UpdateTodoParams struct {
	ID    int32
	Title string
	Task  sql.NullString
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.queryRow(ctx, q.updateTodoStmt, UpdateTodo, arg.ID, arg.Title, arg.Task)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Task,
		&i.Completed,
		&i.DueDate,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
