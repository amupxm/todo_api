package db

import "database/sql"

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	s := Store{
		db:      db,
		Queries: New(db),
	}
	return s
}
