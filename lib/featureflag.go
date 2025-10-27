package featureflag

func ValidateAndParse(xmlPath string) (*Features, error) {
	if err := validateXML(xmlPath); err != nil {
		return nil, err
	}
	if features, err := Parse(xmlPath); err != nil {
		return nil, err
	} else {
		return &features, nil
	}
}
