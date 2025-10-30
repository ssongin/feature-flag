package featureflag

import (
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
	Label   string   `yaml:"label"`
	Value   string   `yaml:"value"`
	Default string   `yaml:"default"`
	Options []string `yaml:"options"`
}

func ParseYAML(data []byte) (Features, error) {
	var root Root
	if err := yaml.Unmarshal(data, &root); err != nil {
		return Features{}, err
	}
	return root.Features, nil
}
