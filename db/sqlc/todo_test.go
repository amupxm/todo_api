package db

import (
	"context"
	"testing"

	"database/sql"

	"github.com/amupxm/todo_api/util"
	"github.com/stretchr/testify/require"
)

var createdTodos []Todo

// ----- Helper Functions ----
func createUser(t *testing.T) User {
	userParams := CreateUserParams{
		Username:  util.RandomName(),
		Password:  util.RandomString(64),
		Email:     util.RandomEmail(),
		FirstName: sql.NullString{String: util.RandomName(), Valid: true},
		LastName:  sql.NullString{String: util.RandomName(), Valid: true},
	}

	user, err := testQuery.CreateUser(context.Background(), userParams)
	require.NoError(t, err)
	return user
}

func createTodo(t *testing.T, user User) Todo {
	todoParams := CreateTodoParams{
		Task:   util.ToSqlNullString(util.RandomName()),
		Title:  util.RandomName(),
		UserID: user.ID,
	}
	todo, err := testQuery.CreateTodo(context.Background(), todoParams)
	require.NoError(t, err)
	return todo
}

// ----- Test Users ----
// -- name: CreateUser :one
func TestCreateUser(t *testing.T) {
	createUser(t)
}

// -- name: UpdateUser :one
func TestUpdateUser(t *testing.T) {
	user := createUser(t)
	arg := UpdateUserParams{
		ID:       user.ID,
		Username: util.RandomName(),
		Email:    util.RandomEmail(),
	}
	user2, err := testQuery.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Email, user2.Email)
}

// -- name: GetUser :one
func TestGetUser(t *testing.T) {
	user := createUser(t)
	user2, err := testQuery.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Email, user2.Email)
}

// -- name: DeleteUser :one
func TestDeleteUser(t *testing.T) {
	user := createUser(t)
	user2, err := testQuery.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.False(t, user.DeletedAt.Valid)
	require.True(t, user2.DeletedAt.Valid)
}

// -- name: CreateTodo :one
func TestCreateTodo(t *testing.T) {
	user := createUser(t)
	createTodo(t, user)
}

// -- name: GetTodosByUser :many
func TestGetTodosByUser(t *testing.T) {
	user := createUser(t)
	for i := 0; i < 10; i++ {
		createTodo(t, user)
	}
	todos, err := testQuery.GetTodosByUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, todos, 10)

}

// -- name: GetTodoByID :one
func TestGetTodoByID(t *testing.T) {
	user := createUser(t)
	todo := createTodo(t, user)
	todo2, err := testQuery.GetTodoByID(context.Background(), todo.ID)
	require.NoError(t, err)
	require.Equal(t, todo, todo2)
}

// -- name: ListTodos :many
func TestListTodos(t *testing.T) {
	user := createUser(t)
	for i := 0; i < 10; i++ {
		createTodo(t, user)
	}
	todos, err := testQuery.ListTodos(context.Background(), ListTodosParams{
		Limit:  10,
		Offset: 1,
	})
	require.NoError(t, err)
	require.Len(t, todos, 10)
}

// -- name: UpdateTodo :one
func TestUpdateTodo(t *testing.T) {
	user := createUser(t)
	todo := createTodo(t, user)
	arg := UpdateTodoParams{
		ID:    todo.ID,
		Title: util.RandomName(),
	}
	todo2, err := testQuery.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Title, todo2.Title)
}

// -- name: DeleteTodo :exec
func TestDeleteTodoByID(t *testing.T) {
	user := createUser(t)
	todo := createTodo(t, user)
	err := testQuery.DeleteTodo(context.Background(), todo.ID)
	require.NoError(t, err)
}

// -- name: ToggleTodo :one
func TestToggleTodoByID(t *testing.T) {
	user := createUser(t)
	todo := createTodo(t, user)
	todo2, err := testQuery.ToggleTodo(
		context.Background(),
		todo.ID,
	)
	require.NoError(t, err)
	require.Equal(t, todo.Completed.Bool, !(todo2.Completed.Bool))
}
