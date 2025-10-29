package featureflag

import (
	_ "embed"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed schema/featureflag.schema.json
var featureFlagSchema []byte

var schemaLoader gojsonschema.JSONLoader
var schema *gojsonschema.Schema

func init() {
	schemaLoader = gojsonschema.NewBytesLoader(featureFlagSchema)
	var err error
	schema, err = gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		panic(fmt.Sprintf("schema compile error: %v", err))
	}
}

func ValidateYAML(yamlBytes []byte) error {
	var doc interface{}
	if err := yaml.Unmarshal(yamlBytes, &doc); err != nil {
		return fmt.Errorf("invalid YAML: %w", err)
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(doc))
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if !result.Valid() {
		var errs []string
		for _, desc := range result.Errors() {
			errs = append(errs, desc.String())
		}
		return fmt.Errorf("validation failed:\n%s", errs)
	}

	return nil
}
