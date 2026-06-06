package schema

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	data := JSON()
	if len(data) == 0 {
		t.Error("JSON() returned empty data")
	}

	// Verify it's valid JSON
	var schema map[string]any
	if err := json.Unmarshal(data, &schema); err != nil {
		t.Errorf("JSON() returned invalid JSON: %v", err)
	}

	// Check expected schema fields
	if schema["$schema"] == nil {
		t.Error("schema missing $schema field")
	}
	if schema["title"] == nil {
		t.Error("schema missing title field")
	}
	if schema["properties"] == nil {
		t.Error("schema missing properties field")
	}
}

func TestString(t *testing.T) {
	s := String()
	if len(s) == 0 {
		t.Error("String() returned empty string")
	}
	if s[0] != '{' {
		t.Error("String() does not start with '{'")
	}
}

func TestNewValidator(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}
	if v == nil {
		t.Error("NewValidator() returned nil")
	}
}

func TestValidateValidJSON(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	validJSON := []byte(`{
		"title": "Test Report",
		"product": {
			"name": "Test Product"
		},
		"author": {
			"email": "test@example.com"
		}
	}`)

	if err := v.Validate(validJSON); err != nil {
		t.Errorf("Validate() error = %v", err)
	}
}

func TestValidateInvalidJSON(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	// Missing required product field
	invalidJSON := []byte(`{
		"title": "Test Report",
		"author": {
			"email": "test@example.com"
		}
	}`)

	if err := v.Validate(invalidJSON); err == nil {
		t.Error("Validate() should return error for invalid JSON")
	}
}

func TestValidateMalformedJSON(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	malformedJSON := []byte(`{not valid json}`)

	if err := v.Validate(malformedJSON); err == nil {
		t.Error("Validate() should return error for malformed JSON")
	}
}

func TestValidateAny(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	validDoc := map[string]any{
		"title": "Test Report",
		"product": map[string]any{
			"name": "Test Product",
		},
		"author": map[string]any{
			"email": "test@example.com",
		},
	}

	if err := v.ValidateAny(validDoc); err != nil {
		t.Errorf("ValidateAny() error = %v", err)
	}
}

func TestValidateAnyInvalid(t *testing.T) {
	v, err := NewValidator()
	if err != nil {
		t.Fatalf("NewValidator() error = %v", err)
	}

	// Missing required fields
	invalidDoc := map[string]any{
		"title": "Test Report",
	}

	if err := v.ValidateAny(invalidDoc); err == nil {
		t.Error("ValidateAny() should return error for invalid document")
	}
}
