package valueobjects

import (
	"fmt"
	"os"

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

func DefinitionFromYAML(data []byte) (Definition, error) {
	def := Definition{}
	err := yaml.Unmarshal(data, &def)
	if err != nil {
		return def, fmt.Errorf("failed reading yaml: %w", err)
	}
	return def, nil
}

func DefinitionFromYAMLFile(path string) (Definition, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Definition{}, fmt.Errorf("failed reading from file: %v", err)
	}
	return DefinitionFromYAML(data)
}
