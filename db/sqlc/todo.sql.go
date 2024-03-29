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
INSERT INTO todos (user_id, title, task)
VALUES ($1, $2, $3)
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
`

type CreateTodoParams struct {
	UserID int32
	Title  string
	Task   sql.NullString
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.queryRow(ctx, q.createTodoStmt, CreateTodo, arg.UserID, arg.Title, arg.Task)
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
		&i.UserID,
	)
	return i, err
}

const CreateUser = `-- name: CreateUser :one
INSERT INTO users (username, password, email, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, password, email, first_name, last_name, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Username  string
	Password  string
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, CreateUser,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const DeleteTodo = `-- name: DeleteTodo :exec
UPDATE todos
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteTodoStmt, DeleteTodo, id)
	return err
}

const DeleteUser = `-- name: DeleteUser :one
UPDATE users
SET deleted_at = now()
WHERE id = $1
RETURNING id, username, password, email, first_name, last_name, created_at, updated_at, deleted_at
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.deleteUserStmt, DeleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const GetTodoByID = `-- name: GetTodoByID :one
SELECT id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
FROM todos
WHERE id = $1
AND  deleted_at IS NULL
LIMIT 1
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
		&i.UserID,
	)
	return i, err
}

const GetTodosByUser = `-- name: GetTodosByUser :many
SELECT id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
FROM todos
WHERE user_id = $1 AND deleted_at IS NULL
`

func (q *Queries) GetTodosByUser(ctx context.Context, userID int32) ([]Todo, error) {
	rows, err := q.query(ctx, q.getTodosByUserStmt, GetTodosByUser, userID)
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
			&i.UserID,
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

const GetUser = `-- name: GetUser :one
SELECT id, username, password, email, first_name, last_name, created_at, updated_at, deleted_at
FROM users
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, GetUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const ListTodos = `-- name: ListTodos :many
SELECT id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
FROM todos
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1 OFFSET $2
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
			&i.UserID,
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
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
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
		&i.UserID,
	)
	return i, err
}

const UpdateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $2,
    task = $3
WHERE id = $1
RETURNING id, title, task, completed, due_date, created_at, updated_at, deleted_at, user_id
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
		&i.UserID,
	)
	return i, err
}

const UpdateUser = `-- name: UpdateUser :one
UPDATE users
SET username = $2,
    password = $3,
    email = $4,
    first_name = $5,
    last_name = $6
WHERE id = $1
RETURNING id, username, password, email, first_name, last_name, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserStmt, UpdateUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
