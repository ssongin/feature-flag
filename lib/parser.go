package featureflag

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Root struct {
	Features Features `yaml:"features"`
}

type Features struct {
	Clusters []Cluster `yaml:"clusters"`
}

type Cluster struct {
	Label        string           `yaml:"label"`
	Clusters     []Cluster        `yaml:"clusters,omitempty"`
	BooleanNodes []BooleanNode    `yaml:"boolean_node,omitempty"`
	PercentNodes []PercentageNode `yaml:"percentage_node,omitempty"`
	StringNodes  []StringNode     `yaml:"string_node,omitempty"`
	ChoiceNodes  []ChoiceNode     `yaml:"choice_node,omitempty"`
}

type BooleanNode struct {
	Label   string `yaml:"label"`
	Value   bool   `yaml:"value"`
	Default bool   `yaml:"default"`
}

type StringNode struct {
	Label   string `yaml:"label"`
	Value   string `yaml:"value"`
	Default string `yaml:"default"`
}

type PercentageNode struct {
	Label   string `yaml:"label"`
	Value   int    `yaml:"value"`
	Default int    `yaml:"default"`
}

type ChoiceNode struct {
	Label   string   `yaml:"label"`
	Value   string   `yaml:"value"`
	Default string   `yaml:"default"`
	Options []string `yaml:"options"`
}

func ParseYAML(path string) (Features, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var root Root
	if err := yaml.Unmarshal(data, &root); err != nil {
		return Features{}, err
	}
	return root.Features, nil
}
