package featureflag

import (
	"fmt"
	"testing"
)

func TestParseYAML_Valid(t *testing.T) {
	yamlData := `
features:
  clusters:
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
            - a
            - b`

	if features, err := ParseYAML([]byte(yamlData)); err != nil {
		t.Fatalf("expected valid YAML, got error: %v", err)
	} else {
		fmt.Println("Result: ")
		fmt.Printf("%#v\n", features)
	}
}

func TestParseYAMLWithMultipleNodes_Valid(t *testing.T) {
	yamlData := `
features:
  clusters:
    - label: "root"
      boolean_node:
        - label: "flag1"
          value: true
          default: false
        - label: "flag2"
          value: false
          default: false
      clusters:
        - label: "child"
          boolean_node:
            - label: "flag3"
              value: true
              default: true
`
	if features, err := ParseYAML([]byte(yamlData)); err != nil {
		t.Fatalf("expected valid YAML, got error: %v", err)
	} else {
		fmt.Printf("%#v\n", features)
	}
}
