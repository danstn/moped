package pipeline

import (
	"log"
)

type API interface {
	Register()
}

type Pipeline struct {
	ID         string
	Name       string
	Definition *Definition
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Register() {
	log.Println("[pipeline] registering pipeline")
}
