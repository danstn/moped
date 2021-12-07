package sqlite

import (
	"database/sql"

	"github.com/danstn/moped/domain/pipeline"
	"github.com/google/uuid"
)

type SQLiteRepositoryConfig interface {
	GetDB() *sql.DB
	GetPipelineTableName() string
}

type SQLiteRepository struct {
	db        *sql.DB
	tableName string
}

func New(config SQLiteRepositoryConfig) *SQLiteRepository {
	return &SQLiteRepository{
		db:        config.GetDB(),
		tableName: config.GetPipelineTableName(),
	}
}

func (r *SQLiteRepository) GetByID(id uuid.UUID) (pipeline.Pipeline, error) {
	panic("not implemented")
}

func (r *SQLiteRepository) Add(p pipeline.Pipeline) (uuid.UUID, error) {
	panic("not implemented")
}
