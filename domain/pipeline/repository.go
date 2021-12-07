package pipeline

import "github.com/google/uuid"

type PipelineRepository interface {
	GetByID(uuid.UUID) (Pipeline, error)
	Add(Pipeline) (uuid.UUID, error)
}
