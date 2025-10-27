package featureflag

import (
	"encoding/xml"
	"os"
)

type Features struct {
	XMLName xml.Name  `xml:"features"`
	Cluster []Cluster `xml:"cluster"`
}

type Cluster struct {
	Label        string           `xml:"label"`
	Clusters     []Cluster        `xml:"cluster"`
	BooleanNodes []BooleanNode    `xml:"boolean_node"`
	PercentNodes []PercentageNode `xml:"percentage_node"`
	StringNodes  []StringNode     `xml:"string_node"`
	ChoiceNodes  []ChoiceNode     `xml:"choice_node"`
}

type BooleanNode struct {
	Label   string `xml:"label"`
	Value   bool   `xml:"value"`
	Default bool   `xml:"default"`
}

type StringNode struct {
	Label   string `xml:"label"`
	Value   string `xml:"value"`
	Default string `xml:"default"`
}

type PercentageNode struct {
	Label   string `xml:"label"`
	Value   int    `xml:"value"`
	Default int    `xml:"default"`
}

type ChoiceNode struct {
	Label   string  `xml:"label"`
	Value   string  `xml:"value"`
	Default string  `xml:"default"`
	Options Options `xml:"options"`
}

type Options struct {
	Option []string `xml:"option"`
}

func Parse(path string) (Features, error) {
	file, err := os.Open(path)
	if err != nil {
		return Features{}, err
	}

	defer file.Close()

	var features Features
	if err := xml.NewDecoder(file).Decode(&features); err != nil {
		return Features{}, err
	}

	return features, nil
}
