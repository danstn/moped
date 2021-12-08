package pipeline

import "github.com/google/uuid"

type Repository interface {
	GetByID(uuid.UUID) (Pipeline, error)
	Save(Pipeline) error
}
