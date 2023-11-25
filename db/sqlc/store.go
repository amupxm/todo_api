package db

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	s := &SQLStore{
		db:      db,
		Queries: New(db),
	}
	return s
}
