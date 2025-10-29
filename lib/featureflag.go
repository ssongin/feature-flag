package featureflag

import (
	"fmt"
	"log"
	"os"
)

func ValidateAndParse(yamlPath string) (*Features, error) {
	yamlData, err := os.ReadFile(yamlPath)
	if err != nil {
		log.Fatalf("failed to read YAML: %v", err)
	}

	if err := ValidateYAML(yamlData); err != nil {
		log.Fatalf("❌ YAML validation failed:\n%v", err)
	}
	fmt.Println("✅ YAML validation passed")

	features, err := ParseYAML(yamlData)
	if err != nil {
		log.Fatalf("❌ Failed to parse YAML: %v", err)
	}
	return &features, nil
}
