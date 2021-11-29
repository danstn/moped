package pipeline

import (
	"log"
)

type API interface {
	Register()
}

type Pipeline struct{}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Register() {
	log.Println("[pipeline] registering pipeline")
}
