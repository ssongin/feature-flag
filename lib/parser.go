package featureflag

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Features struct {
	Clusters []Cluster `yaml:"clusters"`
}

type Cluster struct {
	Label        string           `yaml:"label"`
	Clusters     []Cluster        `yaml:"clusters"`
	BooleanNodes []BooleanNode    `yaml:"boolean_node"`
	PercentNodes []PercentageNode `yaml:"percentage_node"`
	StringNodes  []StringNode     `yaml:"string_node"`
	ChoiceNodes  []ChoiceNode     `yaml:"choice_node"`
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
	Label   string  `yaml:"label"`
	Value   string  `yaml:"value"`
	Default string  `yaml:"default"`
	Options Options `yaml:"options"`
}

type Options struct {
	Option []string `yaml:"option"`
}

func ParseYAML(yamlData []byte) (Features, error) {
	var features Features
	if err := yaml.Unmarshal(yamlData, &features); err != nil {
		return Features{}, fmt.Errorf("failed to parse YAML: %w", err)
	}
	return features, nil
}
