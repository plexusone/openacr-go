package openacr

import (
	"testing"
)

func TestNewReport(t *testing.T) {
	r := NewReport(
		WithProduct("Test Product", "1.0.0"),
		WithAuthor("Test Author", "test@example.com"),
		WithCatalog("2.5-edition-wcag-2.2-508-en"),
	)

	if r.Product.Name != "Test Product" {
		t.Errorf("expected product name 'Test Product', got '%s'", r.Product.Name)
	}
	if r.Product.Version != "1.0.0" {
		t.Errorf("expected product version '1.0.0', got '%s'", r.Product.Version)
	}
	if r.Author.Name != "Test Author" {
		t.Errorf("expected author name 'Test Author', got '%s'", r.Author.Name)
	}
	if r.Author.Email != "test@example.com" {
		t.Errorf("expected author email 'test@example.com', got '%s'", r.Author.Email)
	}
	if r.Catalog != "2.5-edition-wcag-2.2-508-en" {
		t.Errorf("expected catalog '2.5-edition-wcag-2.2-508-en', got '%s'", r.Catalog)
	}
	if r.Chapters == nil {
		t.Error("expected Chapters to be initialized")
	}
}

func TestNewReportWithAllOptions(t *testing.T) {
	r := NewReport(
		WithTitle("Test Report"),
		WithProductDescription("Product", "2.0.0", "A test product"),
		WithAuthorCompany("Jane Doe", "jane@example.com", "Test Inc."),
		WithVendor("John Smith", "john@example.com"),
		WithCatalog("2.5-edition-wcag-2.2-508-en"),
		WithReportDate("2025-01-15"),
		WithVersion(1),
		WithNotes("Test notes"),
		WithEvaluationMethods("Manual testing"),
		WithLegalDisclaimer("Test disclaimer"),
		WithLicense("MIT"),
		WithRepository("https://github.com/example/repo"),
		WithFeedback("https://github.com/example/repo/issues"),
	)

	if r.Title != "Test Report" {
		t.Errorf("expected title 'Test Report', got '%s'", r.Title)
	}
	if r.Product.Description != "A test product" {
		t.Errorf("expected product description 'A test product', got '%s'", r.Product.Description)
	}
	if r.Author.CompanyName != "Test Inc." {
		t.Errorf("expected author company 'Test Inc.', got '%s'", r.Author.CompanyName)
	}
	if r.Vendor == nil {
		t.Fatal("expected vendor to be set")
	}
	if r.Vendor.Name != "John Smith" {
		t.Errorf("expected vendor name 'John Smith', got '%s'", r.Vendor.Name)
	}
	if r.ReportDate != "2025-01-15" {
		t.Errorf("expected report date '2025-01-15', got '%s'", r.ReportDate)
	}
	if r.Version != 1 {
		t.Errorf("expected version 1, got %d", r.Version)
	}
	if r.Notes != "Test notes" {
		t.Errorf("expected notes 'Test notes', got '%s'", r.Notes)
	}
	if r.EvaluationMethods != "Manual testing" {
		t.Errorf("expected evaluation methods 'Manual testing', got '%s'", r.EvaluationMethods)
	}
	if r.LegalDisclaimer != "Test disclaimer" {
		t.Errorf("expected legal disclaimer 'Test disclaimer', got '%s'", r.LegalDisclaimer)
	}
	if r.License != "MIT" {
		t.Errorf("expected license 'MIT', got '%s'", r.License)
	}
	if r.Repository != "https://github.com/example/repo" {
		t.Errorf("expected repository URL, got '%s'", r.Repository)
	}
	if r.Feedback != "https://github.com/example/repo/issues" {
		t.Errorf("expected feedback URL, got '%s'", r.Feedback)
	}
}

func TestComponentNameIsValid(t *testing.T) {
	tests := []struct {
		name     ComponentName
		expected bool
	}{
		{ComponentWeb, true},
		{ComponentElectronicDoc, true},
		{ComponentSoftware, true},
		{ComponentAuthoringTool, true},
		{ComponentNone, true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			if got := tt.name.IsValid(); got != tt.expected {
				t.Errorf("ComponentName(%q).IsValid() = %v, want %v", tt.name, got, tt.expected)
			}
		})
	}
}

func TestAdherenceLevelIsValid(t *testing.T) {
	tests := []struct {
		level    AdherenceLevel
		expected bool
	}{
		{LevelSupports, true},
		{LevelPartiallySupports, true},
		{LevelDoesNotSupport, true},
		{LevelNotApplicable, true},
		{LevelNotEvaluated, true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.level), func(t *testing.T) {
			if got := tt.level.IsValid(); got != tt.expected {
				t.Errorf("AdherenceLevel(%q).IsValid() = %v, want %v", tt.level, got, tt.expected)
			}
		})
	}
}

func TestValidComponentNames(t *testing.T) {
	names := ValidComponentNames()
	if len(names) != 5 {
		t.Errorf("expected 5 component names, got %d", len(names))
	}
}

func TestValidAdherenceLevels(t *testing.T) {
	levels := ValidAdherenceLevels()
	if len(levels) != 5 {
		t.Errorf("expected 5 adherence levels, got %d", len(levels))
	}
}
