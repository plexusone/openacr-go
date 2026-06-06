# openacr-go

Go library for working with OpenACR (Open Accessibility Conformance Report) documents.

OpenACR is a machine-readable format for accessibility conformance reports, based on the VPAT (Voluntary Product Accessibility Template) format. This library supports reading, writing, and validating OpenACR documents in both YAML and JSON formats.

## Installation

```bash
go get github.com/plexusone/openacr-go
```

## Quick Start

### Load an existing report

```go
import "github.com/plexusone/openacr-go"

report, err := openacr.Load("report.yaml")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Product: %s %s\n", report.Product.Name, report.Product.Version)
```

### Create a new report

```go
report := openacr.NewReport(
    openacr.WithProduct("My Product", "1.0.0"),
    openacr.WithAuthor("Jane Doe", "jane@example.com"),
    openacr.WithCatalog("2.5-edition-wcag-2.2-508-en"),
    openacr.WithReportDateNow(),
)

// Add chapter evaluations
report.Chapters["success_criteria_level_a"] = openacr.Chapter{
    Criteria: []openacr.Criterion{
        {
            Num: "1.1.1",
            Components: []openacr.Component{
                {
                    Name: openacr.ComponentWeb,
                    Adherence: openacr.Adherence{
                        Level: openacr.LevelSupports,
                        Notes: "All images have appropriate alt text.",
                    },
                },
            },
        },
    },
}
```

### Validate a report

```go
// Basic validation
errs := report.Validate()
for _, err := range errs {
    fmt.Printf("Validation error: %s\n", err)
}

// Validate against a catalog
cat, _ := catalog.Get("2.5-edition-wcag-2.2-508-en")
errs = report.ValidateAgainstCatalog(cat)
```

### Save a report

```go
// Save as YAML
err := report.Save("report.yaml")

// Save as JSON
err = report.Save("report.json")

// Get bytes
yamlBytes, _ := report.YAML()
jsonBytes, _ := report.JSON()
```

## Catalogs

The catalog subpackage provides access to embedded accessibility standard catalogs from the GSA OpenACR project:

```go
import "github.com/plexusone/openacr-go/catalog"

// List available catalogs
catalogs := catalog.List()
// ["2.5-edition-wcag-2.2-508-en", "2.5-edition-wcag-2.1-508-en", ...]

// Get a catalog
cat, err := catalog.Get("2.5-edition-wcag-2.2-508-en")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Catalog: %s\n", cat.Title)
```

## Schema Validation

The schema subpackage provides JSON Schema validation:

```go
import "github.com/plexusone/openacr-go/schema"

validator, err := schema.NewValidator()
if err != nil {
    log.Fatal(err)
}

jsonData, _ := report.JSON()
if err := validator.Validate(jsonData); err != nil {
    log.Printf("Schema validation failed: %v", err)
}
```

## Adherence Levels

The following adherence levels are supported:

| Level | Description |
|-------|-------------|
| `supports` | The functionality meets the criterion without known defects |
| `partially-supports` | Some functionality does not meet the criterion |
| `does-not-support` | The majority of functionality does not meet the criterion |
| `not-applicable` | The criterion is not relevant to the product |
| `not-evaluated` | The product has not been evaluated (WCAG AAA only) |

## Component Types

The following component types are supported:

| Component | Description |
|-----------|-------------|
| `web` | Web content |
| `electronic-docs` | Electronic documents |
| `software` | Software applications |
| `authoring-tool` | Authoring tools |
| `none` | No specific component |

## Available Catalogs

The following catalogs are embedded:

- `2.5-edition-wcag-2.2-508-en` - WCAG 2.2 with Section 508 (latest)
- `2.5-edition-wcag-2.2-508-eu-en` - WCAG 2.2 with Section 508 and EU standards
- `2.5-edition-wcag-2.2-en` - WCAG 2.2 only
- `2.5-edition-wcag-2.1-508-en` - WCAG 2.1 with Section 508
- `2.5-edition-wcag-2.0-508-en` - WCAG 2.0 with Section 508
- `2.4-edition-wcag-2.1-508-en` - WCAG 2.1 with Section 508 (2.4 edition)
- `2.4-edition-wcag-2.1-508-eu-en` - WCAG 2.1 with Section 508 and EU (2.4 edition)
- `2.4-edition-wcag-2.1-en` - WCAG 2.1 only (2.4 edition)
- `2.4-edition-wcag-2.0-508-en` - WCAG 2.0 with Section 508 (2.4 edition)

## License

MIT License - see LICENSE file.

## References

- [GSA OpenACR](https://github.com/GSA/openacr) - Original OpenACR specification
- [VPAT](https://www.itic.org/policy/accessibility/vpat) - Voluntary Product Accessibility Template
- [WCAG](https://www.w3.org/WAI/standards-guidelines/wcag/) - Web Content Accessibility Guidelines
