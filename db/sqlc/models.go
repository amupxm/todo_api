// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
)

type Todo struct {
	ID        int32
	Title     string
	Task      sql.NullString
	Completed sql.NullBool
	DueDate   sql.NullTime
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
	UserID    int32
}

type User struct {
	ID        int32
	Username  string
	Password  string
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
