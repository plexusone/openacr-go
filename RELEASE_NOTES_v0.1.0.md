# Release Notes: v0.1.0

**Release Date:** 2026-06-06

## Overview

Initial release of `openacr-go`, a Go library for working with [OpenACR](https://github.com/GSA/openacr) (Open Accessibility Conformance Report) documents.

OpenACR is a machine-readable format for accessibility conformance reports developed by the GSA, based on the VPAT (Voluntary Product Accessibility Template) format. This library provides full support for reading, writing, and validating OpenACR documents in both YAML and JSON formats.

## Features

### Core Types

- `Report` - Root type representing an OpenACR document
- `Product` - Product information (name, version, description)
- `Contact` - Author/vendor contact information
- `Chapter` - WCAG chapter evaluations
- `Criterion` - Individual success criterion assessment
- `Component` - Component-level adherence (web, software, electronic-docs, authoring-tool)
- `Adherence` - Conformance level and notes

### Loading and Writing

- Load OpenACR files from disk with automatic format detection
- Read from `io.Reader` in YAML or JSON format
- Write to disk or `io.Writer` in either format
- Get raw bytes via `YAML()` and `JSON()` methods

### Functional Options

Create reports with a clean, fluent API:

```go
report := openacr.NewReport(
    openacr.WithProduct("My App", "1.0.0"),
    openacr.WithAuthor("Jane Doe", "jane@example.com"),
    openacr.WithCatalog("2.5-edition-wcag-2.2-508-en"),
    openacr.WithReportDateNow(),
)
```

### Validation

- `Validate()` - Basic structural validation
- `ValidateAgainstCatalog()` - Validate criteria against catalog definitions

### Embedded Catalogs

Nine GSA catalogs are embedded via `go:embed`:

| Catalog | Description |
|---------|-------------|
| `2.5-edition-wcag-2.2-508-en` | WCAG 2.2 + Section 508 (recommended) |
| `2.5-edition-wcag-2.2-508-eu-en` | WCAG 2.2 + Section 508 + EU |
| `2.5-edition-wcag-2.2-en` | WCAG 2.2 only |
| `2.5-edition-wcag-2.1-508-en` | WCAG 2.1 + Section 508 |
| `2.5-edition-wcag-2.0-508-en` | WCAG 2.0 + Section 508 |
| `2.4-edition-wcag-2.1-508-en` | WCAG 2.1 + Section 508 (2.4 edition) |
| `2.4-edition-wcag-2.1-508-eu-en` | WCAG 2.1 + Section 508 + EU (2.4 edition) |
| `2.4-edition-wcag-2.1-en` | WCAG 2.1 only (2.4 edition) |
| `2.4-edition-wcag-2.0-508-en` | WCAG 2.0 + Section 508 (2.4 edition) |

### JSON Schema Validation

The `schema` package provides JSON Schema validation using the embedded `openacr-0.1.0.json` schema:

```go
validator, _ := schema.NewValidator()
err := validator.Validate(jsonData)
```

## Installation

```bash
go get github.com/plexusone/openacr-go@v0.1.0
```

## Quick Start

```go
import "github.com/plexusone/openacr-go"

// Load existing report
report, _ := openacr.Load("report.yaml")

// Create new report
report := openacr.NewReport(
    openacr.WithProduct("My App", "1.0.0"),
    openacr.WithAuthor("A11y Team", "a11y@example.com"),
    openacr.WithCatalog("2.5-edition-wcag-2.2-508-en"),
)

// Validate
errs := report.Validate()

// Save
report.Save("output.yaml")
```

## Dependencies

- `gopkg.in/yaml.v3` - YAML parsing
- `github.com/santhosh-tekuri/jsonschema/v5` - JSON Schema validation

## Links

- [GitHub Repository](https://github.com/plexusone/openacr-go)
- [GSA OpenACR Specification](https://github.com/GSA/openacr)
- [VPAT](https://www.itic.org/policy/accessibility/vpat)
- [WCAG](https://www.w3.org/WAI/standards-guidelines/wcag/)
