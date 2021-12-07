package pipeline

import (
	"errors"

	"github.com/danstn/moped/entity"
	"github.com/danstn/moped/valueobjects"
	"github.com/google/uuid"
)

var (
	ErrEmptyPipelineName = errors.New("name cannot be empty")
)

type Pipeline struct {
	pipeline   *entity.Pipeline
	definition valueobjects.Definition
}

func NewPipeline(name string, definition valueobjects.Definition) (Pipeline, error) {
	if name == "" {
		return Pipeline{}, ErrEmptyPipelineName
	}

	pipeline := &entity.Pipeline{
		ID:   uuid.New(),
		Name: name,
	}

	return Pipeline{
		pipeline:   pipeline,
		definition: definition,
	}, nil
}
