package config

import "github.com/danstn/moped/internal/pipeline"

type AppConfig struct {
	pipeline *pipeline.Pipeline
}

func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{
		pipeline: pipeline.NewPipeline(),
	}
	return config, nil
}

func (c *AppConfig) GetPipeline() pipeline.API {
	return c.pipeline
}
