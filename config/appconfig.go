package config

import (
	"fmt"

	pipelineSQLiteStore "github.com/danstn/moped/domain/pipeline/sqlite"
	"github.com/danstn/moped/pkg/sqlite"
	pipelineService "github.com/danstn/moped/services/pipeline"
	_ "github.com/mattn/go-sqlite3"
)

type AppConfig struct {
	pipelineService pipelineService.API
}

func NewAppConfig() (*AppConfig, error) {
	// open db
	db, err := sqlite.Open("./dev.db")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// create repositories
	pipelineRepo, err := pipelineSQLiteStore.New(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline repo: %w", err)
	}

	// create services
	pipelineSvc, err := pipelineService.New(
		pipelineService.WithSQLiteRepository(pipelineRepo),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create a pipeline service: %v", err)
	}

	// create config
	return &AppConfig{
		pipelineService: pipelineSvc,
	}, nil
}

// open db connection

// resource accessorts

func (c *AppConfig) GetPipelineService() pipelineService.API {
	return c.pipelineService
}
