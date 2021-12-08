package pipeline

import (
	"errors"

	entity "github.com/danstn/moped/entity"
	valueobjects "github.com/danstn/moped/valueobjects"
	"github.com/google/uuid"
)

var (
	ErrEmptyPipelineName = errors.New("name cannot be empty")
)

type Pipeline struct {
	pipeline   *entity.Pipeline
	definition valueobjects.Definition
}

func New(name string, definition valueobjects.Definition) (Pipeline, error) {
	if name == "" {
		return Pipeline{}, ErrEmptyPipelineName
	}

	pipeline := &entity.Pipeline{
		ID:     uuid.New(),
		Name:   name,
		Status: "DISABLED",
	}

	return Pipeline{
		pipeline:   pipeline,
		definition: definition,
	}, nil
}

func (p Pipeline) GetID() uuid.UUID {
	return p.pipeline.ID
}

func (p Pipeline) GetDefinition() valueobjects.Definition {
	return p.definition
}

func (p Pipeline) GetStatus() string {
	return p.pipeline.Status
}
