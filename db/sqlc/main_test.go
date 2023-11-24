package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testQuery *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1234@localhost:5432/todosdb?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQuery = New(conn)
	exitCode := m.Run()

	// Run additional test after all other tests

	// Exit with the combined exit code
	os.Exit(exitCode)
}
func TestToggleTodo(t *testing.T) {
	todo := createTodoInDB(t)
	todo2, err := testQuery.ToggleTodo(
		context.Background(),
		todo.ID,
	)
	require.NoError(t, err)
	require.Equal(t, todo.Completed.Bool, !(todo2.Completed.Bool))
}
