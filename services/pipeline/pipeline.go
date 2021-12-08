package pipeline

import (
	"context"
	"fmt"
	"log"

	"github.com/danstn/moped/domain/pipeline"
	"github.com/danstn/moped/valueobjects"
	"github.com/google/uuid"
)

type API interface {
	CreatePipelineFromFile(context.Context, string) (uuid.UUID, error)
	CreatePipeline(context.Context, valueobjects.Definition) (uuid.UUID, error)
	Schedule(pipelineID uuid.UUID) error
}

type service struct {
	name          string
	pipelinesRepo pipeline.Repository
}

type Configuration func(*service) error

func New(cfgs ...Configuration) (API, error) {
	s := &service{
		name: "pipeline-service",
	}
	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func WithSQLiteRepository(r pipeline.Repository) Configuration {
	return func(s *service) error {
		s.pipelinesRepo = r
		return nil
	}
}

func (s *service) CreatePipelineFromFile(ctx context.Context, file string) (uuid.UUID, error) {
	log.Println("reading pipeline from file:", file)
	definition, err := valueobjects.DefinitionFromYAMLFile(file)
	if err != nil {
		log.Fatalf(s.formatError("failed reading pipeline fie", err).Error())
	}
	return s.CreatePipeline(ctx, definition)
}

func (s *service) CreatePipeline(_ context.Context, definition valueobjects.Definition) (uuid.UUID, error) {
	log.Printf("[svc] creating new pipeline from definition: %v", definition.Name)
	newPipeline, err := pipeline.New(definition.Name, definition)
	if err != nil {
		return uuid.Nil, s.formatError("failed creating new pipeline", err)
	}
	err = s.pipelinesRepo.Save(newPipeline)
	if err != nil {
		return uuid.Nil, s.formatError("failed saving pipeline to store", err)
	}
	return newPipeline.GetID(), nil
}

func (s *service) Schedule(pipelineID uuid.UUID) error {
	log.Printf("scheduling job for pipeline: %v", pipelineID)
	return nil
}

func (s *service) formatError(msg string, e error) error {
	return fmt.Errorf("[%s] %s: %v", s.name, msg, e)
}
