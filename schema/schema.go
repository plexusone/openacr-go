// Package schema provides JSON Schema validation for OpenACR documents.
//
// The JSON Schema is embedded from the GSA OpenACR project and can be used
// to validate OpenACR documents before processing.
//
// # Usage
//
//	validator, err := schema.NewValidator()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	data, _ := os.ReadFile("report.json")
//	if err := validator.Validate(data); err != nil {
//	    log.Printf("Validation failed: %v", err)
//	}
package schema

// JSON returns the embedded OpenACR JSON Schema as bytes.
func JSON() []byte {
	result := make([]byte, len(schemaJSON))
	copy(result, schemaJSON)
	return result
}

// String returns the embedded OpenACR JSON Schema as a string.
func String() string {
	return string(schemaJSON)
}
