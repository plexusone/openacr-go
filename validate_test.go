package openacr

import (
	"errors"
	"testing"

	"github.com/plexusone/openacr-go/catalog"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		report      *Report
		expectErrs  int
		checkFields []string
	}{
		{
			name: "valid report",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
			},
			expectErrs: 0,
		},
		{
			name: "missing product name",
			report: &Report{
				Author: Contact{Email: "test@example.com"},
			},
			expectErrs:  1,
			checkFields: []string{"product.name"},
		},
		{
			name: "missing author email",
			report: &Report{
				Product: Product{Name: "Test"},
			},
			expectErrs:  1,
			checkFields: []string{"author.email"},
		},
		{
			name:       "missing both required fields",
			report:     &Report{},
			expectErrs: 2,
		},
		{
			name: "invalid component name",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"test": {
						Criteria: []Criterion{
							{
								Num: "1.1.1",
								Components: []Component{
									{Name: "invalid-component"},
								},
							},
						},
					},
				},
			},
			expectErrs: 1,
		},
		{
			name: "invalid adherence level",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"test": {
						Criteria: []Criterion{
							{
								Num: "1.1.1",
								Components: []Component{
									{
										Name:      ComponentWeb,
										Adherence: Adherence{Level: "invalid-level"},
									},
								},
							},
						},
					},
				},
			},
			expectErrs: 1,
		},
		{
			name: "valid with chapters",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"success_criteria_level_a": {
						Criteria: []Criterion{
							{
								Num: "1.1.1",
								Components: []Component{
									{
										Name:      ComponentWeb,
										Adherence: Adherence{Level: LevelSupports},
									},
								},
							},
						},
					},
				},
			},
			expectErrs: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.report.Validate()
			if len(errs) != tt.expectErrs {
				t.Errorf("Validate() returned %d errors, expected %d", len(errs), tt.expectErrs)
				for _, err := range errs {
					t.Logf("  error: %s", err)
				}
			}

			// Check for expected fields in errors
			for _, field := range tt.checkFields {
				found := false
				for _, err := range errs {
					if err.Field == field {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("expected error for field %q", field)
				}
			}
		})
	}
}

func TestValidateAgainstCatalog(t *testing.T) {
	cat, err := catalog.Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		t.Fatalf("Failed to load catalog: %v", err)
	}

	tests := []struct {
		name       string
		report     *Report
		expectErrs int
	}{
		{
			name: "valid against catalog",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"success_criteria_level_a": {
						Criteria: []Criterion{
							{
								Num: "1.1.1",
								Components: []Component{
									{
										Name:      ComponentWeb,
										Adherence: Adherence{Level: LevelSupports},
									},
								},
							},
						},
					},
				},
			},
			expectErrs: 0,
		},
		{
			name: "invalid chapter",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"nonexistent_chapter": {
						Criteria: []Criterion{
							{Num: "1.1.1"},
						},
					},
				},
			},
			expectErrs: 1,
		},
		{
			name: "invalid criterion in chapter",
			report: &Report{
				Product: Product{Name: "Test"},
				Author:  Contact{Email: "test@example.com"},
				Chapters: map[string]Chapter{
					"success_criteria_level_a": {
						Criteria: []Criterion{
							{Num: "99.99.99"}, // Invalid criterion
						},
					},
				},
			},
			expectErrs: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.report.ValidateAgainstCatalog(cat)
			if len(errs) != tt.expectErrs {
				t.Errorf("ValidateAgainstCatalog() returned %d errors, expected %d", len(errs), tt.expectErrs)
				for _, err := range errs {
					t.Logf("  error: %s", err)
				}
			}
		})
	}
}

func TestValidationErrorUnwrap(t *testing.T) {
	ve := ValidationError{
		Field:   "test.field",
		Message: "test message",
		Err:     ErrEmptyProductName,
	}

	if !errors.Is(ve, ErrEmptyProductName) {
		t.Error("ValidationError.Unwrap() should return underlying error")
	}

	errStr := ve.Error()
	if errStr != "test.field: test message" {
		t.Errorf("ValidationError.Error() = %q, want %q", errStr, "test.field: test message")
	}
}

func TestValidationErrorWithoutField(t *testing.T) {
	ve := ValidationError{
		Message: "test message",
	}

	errStr := ve.Error()
	if errStr != "test message" {
		t.Errorf("ValidationError.Error() = %q, want %q", errStr, "test message")
	}
}

func TestLoadedReportValidation(t *testing.T) {
	// Test validation of loaded valid.yaml
	report, err := Load("testdata/valid.yaml")
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	errs := report.Validate()
	if len(errs) != 0 {
		t.Errorf("valid.yaml should have no validation errors, got %d", len(errs))
		for _, err := range errs {
			t.Logf("  error: %s", err)
		}
	}
}
