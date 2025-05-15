package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open() (*sql.DB, error) {
	url := "postgres://postgres:postgres@localhost:5432/postgres"
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	fmt.Println("Connected to Database...")

	return db, nil
}
