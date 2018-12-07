package db

import (
	"database/sql"
	// we only support postgres as of now
	_ "github.com/lib/pq"
)

// Env holds the db pool
type Env struct {
	DB *sql.DB
}

// NewEnv returns a new Env instance with db in it
func NewEnv(dsn string) (*Env, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &Env{db}, nil
}
