package db

import (
	"context"
	"testing"

	"database/sql"

	"github.com/amupxm/todo_api/util"
	"github.com/stretchr/testify/require"
)

var createdTodos []Todo

func createTodoInDB(t *testing.T) Todo {
	arg := CreateTodoParams{
		Task:  sql.NullString{String: util.RandomString(), Valid: true},
		Title: util.RandomString(),
	}

	todo, err := testQuery.CreateTodo(context.Background(), arg)

	require.Nil(t, err)
	require.NotEmpty(t, todo.ID)
	require.Equal(t, arg.Task, todo.Task)
	require.False(t, todo.Completed.Bool)
	require.Equal(t, arg.Title, todo.Title)

	createdTodos = append(createdTodos, todo)
	return todo
}

func deleteTodoInDB(t *testing.T, ID int32) {
	err := testQuery.DeleteTodo(context.Background(), ID)

	require.NoError(t, err)
}

func TestCreateTodo(t *testing.T) {
	createTodoInDB(t)
}

func TestDeleteTodo(t *testing.T) {
	todo := createTodoInDB(t)
	deleteTodoInDB(t, todo.ID)
}

func TestGetTodoByID(t *testing.T) {
	todo := createTodoInDB(t)
	todo2, err := testQuery.GetTodoByID(context.Background(), todo.ID)
	require.NoError(t, err)
	require.Equal(t, todo, todo2)
}

func TestListTodos(t *testing.T) {
	for i := 0; i < 10; i++ {
		createTodoInDB(t)
	}
	todos, err := testQuery.ListTodos(context.Background(), ListTodosParams{
		Limit:  10,
		Offset: 1,
	})
	require.NoError(t, err)
	require.Len(t, todos, 10)
}

func TestUpdateTodo(t *testing.T) {
	todo := createTodoInDB(t)
	arg := UpdateTodoParams{
		ID:    todo.ID,
		Title: util.RandomString(),
	}
	todo2, err := testQuery.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Title, todo2.Title)
}
