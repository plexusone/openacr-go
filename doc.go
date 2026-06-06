// Package openacr provides types and utilities for working with OpenACR
// (Open Accessibility Conformance Report) documents.
//
// OpenACR is a machine-readable format for accessibility conformance reports,
// based on the VPAT (Voluntary Product Accessibility Template) format.
// This package supports reading, writing, and validating OpenACR documents
// in both YAML and JSON formats.
//
// # Basic Usage
//
// Load an existing OpenACR report:
//
//	report, err := openacr.Load("report.yaml")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Product: %s %s\n", report.Product.Name, report.Product.Version)
//
// Create a new report:
//
//	report := openacr.NewReport(
//	    openacr.WithProduct("My Product", "1.0.0"),
//	    openacr.WithAuthor("Jane Doe", "jane@example.com"),
//	    openacr.WithCatalog("2.5-edition-wcag-2.2-508-en"),
//	)
//
// Validate a report:
//
//	errors := report.Validate()
//	for _, err := range errors {
//	    fmt.Printf("Validation error: %s\n", err)
//	}
//
// # Catalogs
//
// The catalog subpackage provides access to embedded accessibility standard
// catalogs from the GSA OpenACR project:
//
//	cat, err := catalog.Get("2.5-edition-wcag-2.2-508-en")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Catalog: %s\n", cat.Title)
//
// # Schema Validation
//
// The schema subpackage provides JSON Schema validation:
//
//	validator, err := schema.NewValidator()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	if err := validator.Validate(jsonData); err != nil {
//	    log.Printf("Schema validation failed: %v", err)
//	}
package openacr
