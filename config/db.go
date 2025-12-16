package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewDB(dbURL string) (*sql.DB, error) {
	return sql.Open("postgres", dbURL)
}