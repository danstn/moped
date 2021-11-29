package pipeline

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Definition struct {
	Name  string `json:"name" yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name string `yaml:"name"`
	Run  string `yaml:"run"`
}

func FromYAML(data []byte) (*Definition, error) {
	def := &Definition{}
	err := yaml.Unmarshal(data, def)
	if err != nil {
		return nil, fmt.Errorf("failed reading yaml: %w", err)
	}
	return def, nil
}
