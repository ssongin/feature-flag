package featureflag

import (
	_ "embed"
	"fmt"
	"os"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

//go:embed schema/featureflag.xsd
var featureFlagXSD []byte

func ValidateXML(xmlPath string, xsdPath string) error {
	xmlBytes, err := os.ReadFile(xmlPath)
	if err != nil {
		return fmt.Errorf("failed to read XML: %w", err)
	}

	handler, err := xsdvalidate.NewXsdHandlerMem(featureFlagXSD, xsdvalidate.ParsErrDefault)
	if err != nil {
		return fmt.Errorf("failed to create XSD handler: %w", err)
	}
	defer handler.Free()

	if err := handler.ValidateMem(xmlBytes, xsdvalidate.ValidErrDefault); err != nil {
		return fmt.Errorf("XML validation failed: %w", err)
	}

	return nil
}
