// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createTodoStmt, err = db.PrepareContext(ctx, CreateTodo); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTodo: %w", err)
	}
	if q.deleteTodoStmt, err = db.PrepareContext(ctx, DeleteTodo); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTodo: %w", err)
	}
	if q.getTodoByIDStmt, err = db.PrepareContext(ctx, GetTodoByID); err != nil {
		return nil, fmt.Errorf("error preparing query GetTodoByID: %w", err)
	}
	if q.listTodosStmt, err = db.PrepareContext(ctx, ListTodos); err != nil {
		return nil, fmt.Errorf("error preparing query ListTodos: %w", err)
	}
	if q.toggleTodoStmt, err = db.PrepareContext(ctx, ToggleTodo); err != nil {
		return nil, fmt.Errorf("error preparing query ToggleTodo: %w", err)
	}
	if q.updateTodoStmt, err = db.PrepareContext(ctx, UpdateTodo); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTodo: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createTodoStmt != nil {
		if cerr := q.createTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTodoStmt: %w", cerr)
		}
	}
	if q.deleteTodoStmt != nil {
		if cerr := q.deleteTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTodoStmt: %w", cerr)
		}
	}
	if q.getTodoByIDStmt != nil {
		if cerr := q.getTodoByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTodoByIDStmt: %w", cerr)
		}
	}
	if q.listTodosStmt != nil {
		if cerr := q.listTodosStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTodosStmt: %w", cerr)
		}
	}
	if q.toggleTodoStmt != nil {
		if cerr := q.toggleTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing toggleTodoStmt: %w", cerr)
		}
	}
	if q.updateTodoStmt != nil {
		if cerr := q.updateTodoStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTodoStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db              DBTX
	tx              *sql.Tx
	createTodoStmt  *sql.Stmt
	deleteTodoStmt  *sql.Stmt
	getTodoByIDStmt *sql.Stmt
	listTodosStmt   *sql.Stmt
	toggleTodoStmt  *sql.Stmt
	updateTodoStmt  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:              tx,
		tx:              tx,
		createTodoStmt:  q.createTodoStmt,
		deleteTodoStmt:  q.deleteTodoStmt,
		getTodoByIDStmt: q.getTodoByIDStmt,
		listTodosStmt:   q.listTodosStmt,
		toggleTodoStmt:  q.toggleTodoStmt,
		updateTodoStmt:  q.updateTodoStmt,
	}
}
