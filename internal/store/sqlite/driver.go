package sqlite

import (
	"context"
	"database/sql"

	"github.com/danstn/moped/internal/store"
)

type Pipeline struct {
	db        *sql.DB
	tableName string
}

type pipelineConfig interface {
	store.SQLiteStore
	Insert(context.Context)
}
