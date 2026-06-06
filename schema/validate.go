package schema

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

// Validator validates JSON data against the OpenACR schema.
type Validator struct {
	schema *jsonschema.Schema
}

// NewValidator creates a new Validator with the embedded OpenACR schema.
func NewValidator() (*Validator, error) {
	compiler := jsonschema.NewCompiler()
	compiler.Draft = jsonschema.Draft2020

	if err := compiler.AddResource("openacr.json", bytes.NewReader(schemaJSON)); err != nil {
		return nil, fmt.Errorf("adding schema resource: %w", err)
	}

	schema, err := compiler.Compile("openacr.json")
	if err != nil {
		return nil, fmt.Errorf("compiling schema: %w", err)
	}

	return &Validator{schema: schema}, nil
}

// Validate validates JSON data against the OpenACR schema.
func (v *Validator) Validate(data []byte) error {
	var doc any
	if err := json.Unmarshal(data, &doc); err != nil {
		return fmt.Errorf("decoding JSON: %w", err)
	}

	if err := v.schema.Validate(doc); err != nil {
		return fmt.Errorf("schema validation failed: %w", err)
	}

	return nil
}

// ValidateAny validates any value against the OpenACR schema.
// The value should be a map or struct that can be validated against the schema.
func (v *Validator) ValidateAny(doc any) error {
	if err := v.schema.Validate(doc); err != nil {
		return fmt.Errorf("schema validation failed: %w", err)
	}
	return nil
}
