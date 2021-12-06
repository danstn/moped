package config

import (
	"database/sql"
	"fmt"

	"github.com/danstn/moped/internal/pipeline"
	_ "github.com/mattn/go-sqlite3"
)

type AppConfig struct {
	db       *sql.DB
	pipeline pipeline.API
}

func NewAppConfig() (*AppConfig, error) {
	appConfig := &AppConfig{}

	db, err := newSQLiteClient("./dev.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	appConfig.db = db

	appConfig.pipeline = pipeline.NewService(appConfig)
	return appConfig, nil
}

// open db connection

func newSQLiteClient(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database %v: %w", dbPath, err)
	}
	return db, nil
}

// resource accessorts

func (c *AppConfig) GetPipeline() pipeline.API {
	return c.pipeline
}

func (c *AppConfig) GetSQLite() *sql.DB {
	return c.db
}
