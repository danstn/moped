package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/danstn/moped/internal/pipeline"
	"github.com/google/uuid"
)

type API interface {
	CreatePipeline(context.Context, *pipeline.Definition) (string, error)
}

type StoreConfig interface {
	GetSQLite() *sql.DB
}

type sqliteStore struct {
	db *sql.DB
}

func NewStore(config StoreConfig) API {
	return &sqliteStore{
		db: config.GetSQLite(),
	}
}

func (s *sqliteStore) CreatePipeline(ctx context.Context, pipeline *pipeline.Definition) (string, error) {
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
