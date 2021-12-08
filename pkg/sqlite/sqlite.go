package sqlite

import (
	"database/sql"
	"fmt"
)

func Open(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database %v: %w", dbPath, err)
	}
	return db, nil
}
