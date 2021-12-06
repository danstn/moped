package pipeline

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type API interface {
	CreatePipeline(context.Context, *Definition) (string, error)
}

type Config interface {
	GetSQLite() *sql.DB
}

type sqliteStore struct {
	db *sql.DB
}

func NewService(config Config) API {
	db := config.GetSQLite()
	store := &sqliteStore{
		db: db,
	}
	if err := store.migrate(context.Background()); err != nil {
		log.Fatalf("failed creating new service: %v", err)
	}

	return store
}

func (s *sqliteStore) CreatePipeline(ctx context.Context, pipeline *Definition) (string, error) {
	// marshal pipeline
	json, err := json.Marshal(pipeline)
	if err != nil {
		return "", fmt.Errorf("failed marshaling pipeline: %w", err)
	}
	// generate new ID
	newID := uuid.NewString()
	// insert to DB
	stmt := "INSERT INTO pipeline(id, name, definition) VALUES (?, ?, ?)"
	_, err = s.db.ExecContext(ctx, stmt, newID, pipeline.Name, json)
	if err != nil {
		return "", fmt.Errorf("failed inserting into db: %w", err)
	}
	return newID, nil
}

func (s *sqliteStore) migrate(ctx context.Context) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS "pipeline" (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			definition TEXT NOT NULL
	)`,
	}
	for i, stmt := range migrations {
		log.Printf("[%d] executing migration: \n\t%s", i, stmt)
		_, err := s.db.ExecContext(ctx, stmt)
		if err != nil {
			return fmt.Errorf("db migration failed: %w", err)
		}
	}
	return nil
}
