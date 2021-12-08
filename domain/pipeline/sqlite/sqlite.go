package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/danstn/moped/domain/pipeline"
	"github.com/google/uuid"
)

type SQLiteRepository struct {
	db *sql.DB
}

func New(db *sql.DB) (*SQLiteRepository, error) {
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed migrating db: %v", err)
	}
	return &SQLiteRepository{
		db: db,
	}, nil
}

func (r *SQLiteRepository) GetByID(id uuid.UUID) (pipeline.Pipeline, error) {
	panic("not implemented")
}

func (r *SQLiteRepository) Save(p pipeline.Pipeline) error {
	definition := p.GetDefinition()
	json, err := json.Marshal(definition)
	if err != nil {
		return fmt.Errorf("failed marshaling pipeline: %w", err)
	}
	stmt := "INSERT INTO pipeline(id, name, status, definition) VALUES (?, ?, ?, ?)"
	_, err = r.db.Exec(stmt, p.GetID(), definition.Name, p.GetStatus(), json)
	if err != nil {
		return fmt.Errorf("failed inserting into db: %w", err)
	}
	return nil
}

func migrate(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS "pipeline" (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			definition TEXT NOT NULL,
			status TEXT NOT NULL
	)`,
	}
	for i, stmt := range migrations {
		log.Printf("[%d] executing migration: \n\t%s", i, stmt)
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("db migration failed: %w", err)
		}
	}
	return nil
}
