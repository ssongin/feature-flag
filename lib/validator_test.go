package featureflag

import (
	"testing"
)

func TestValidateYAML_Valid(t *testing.T) {
	yamlData := `
features:
  cluster:
    - label: "root"
      boolean_node:
        - label: "flag1"
          value: true
          default: false
      percentage_node:
        - label: "percent1"
          value: 50
          default: 0
      choice_node:
        - label: "choice1"
          value: "a"
          default: "b"
          options:
            option: ["a", "b"]
`
	if err := ValidateYAML([]byte(yamlData)); err != nil {
		t.Fatalf("expected valid YAML, got error: %v", err)
	}
}

func TestValidateYAMLWithMultipleNodes_Valid(t *testing.T) {
	yamlData := `
features:
  cluster:
    - label: "root"
      boolean_node:
        - label: "flag1"
          value: true
          default: false
        - label: "flag2"
          value: false
          default: false
      cluster:
        - label: "child"
          boolean_node:
            - label: "flag3"
              value: true
              default: true
`
	if err := ValidateYAML([]byte(yamlData)); err != nil {
		t.Fatalf("expected valid YAML, got error: %v", err)
	}
}

func TestValidateYAML_MissingLabel(t *testing.T) {
	yamlData := `
features:
  cluster:
    - boolean_node:
        - label: "flag1"
          value: true
          default: false
`
	if err := ValidateYAML([]byte(yamlData)); err == nil {
		t.Fatal("expected validation error for missing label, got nil")
	}
}

func TestValidateYAML_InvalidType(t *testing.T) {
	yamlData := `
features:
  cluster:
    - label: "root"
      boolean_node:
        - label: "flag1"
          value: "notbool"
          default: false
`
	if err := ValidateYAML([]byte(yamlData)); err == nil {
		t.Fatal("expected validation error for invalid type, got nil")
	}
}

func TestValidateYAML_PercentageOutOfRange(t *testing.T) {
	yamlData := `
features:
  cluster:
    - label: "root"
      percentage_node:
        - label: "p1"
          value: 200
          default: 0
`
	if err := ValidateYAML([]byte(yamlData)); err == nil {
		t.Fatal("expected validation error for out-of-range percentage, got nil")
	}
}

func TestValidateYAML_TooFewOptions(t *testing.T) {
	yamlData := `
features:
  cluster:
    - label: "root"
      choice_node:
        - label: "c1"
          value: "a"
          default: "b"
          options:
            option: ["onlyone"]
`
	if err := ValidateYAML([]byte(yamlData)); err == nil {
		t.Fatal("expected validation error for too few options, got nil")
	}
}

func TestValidateYAML_EmptyCluster(t *testing.T) {
	yamlData := `
features:
  cluster: []
`
	if err := ValidateYAML([]byte(yamlData)); err == nil {
		t.Fatal("expected validation error for empty cluster array, got nil")
	}
}
