package pipeline

import (
	"log"
)

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
