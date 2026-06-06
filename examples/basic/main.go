// Example demonstrates basic usage of the openacr package.
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/plexusone/openacr-go"
	"github.com/plexusone/openacr-go/catalog"
)

func main() {
	// Create a new report with functional options
	report := openacr.NewReport(
		openacr.WithTitle("Example Product Accessibility Conformance Report"),
		openacr.WithProduct("Example Product", "1.0.0"),
		openacr.WithAuthor("Jane Doe", "jane@example.com"),
		openacr.WithCatalog("2.5-edition-wcag-2.2-508-en"),
		openacr.WithReportDateNow(),
		openacr.WithNotes("This is an example accessibility conformance report."),
		openacr.WithEvaluationMethods("Manual testing with NVDA and VoiceOver. Automated testing with axe-core."),
	)

	// Add a chapter with criteria
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
					{
						Name: openacr.ComponentElectronicDoc,
						Adherence: openacr.Adherence{
							Level: openacr.LevelNotApplicable,
						},
					},
				},
			},
			{
				Num: "1.2.1",
				Components: []openacr.Component{
					{
						Name: openacr.ComponentWeb,
						Adherence: openacr.Adherence{
							Level: openacr.LevelPartiallySupports,
							Notes: "Some pre-recorded audio content lacks text alternatives.",
						},
					},
				},
			},
		},
	}

	// Validate the report
	fmt.Println("=== Basic Validation ===")
	if errs := report.Validate(); len(errs) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errs {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("Report is valid!")
	}

	// Validate against catalog
	fmt.Println("\n=== Catalog Validation ===")
	cat, err := catalog.Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		log.Fatalf("Failed to load catalog: %v", err)
	}

	if errs := report.ValidateAgainstCatalog(cat); len(errs) > 0 {
		fmt.Println("Catalog validation errors:")
		for _, err := range errs {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("Report is valid against catalog!")
	}

	// Output as JSON
	fmt.Println("\n=== Report as JSON ===")
	jsonBytes, err := report.JSON()
	if err != nil {
		log.Fatalf("Failed to convert to JSON: %v", err)
	}

	var pretty map[string]any
	if err := json.Unmarshal(jsonBytes, &pretty); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	prettyJSON, _ := json.MarshalIndent(pretty, "", "  ")
	fmt.Println(string(prettyJSON))

	// List available catalogs
	fmt.Println("\n=== Available Catalogs ===")
	for _, id := range catalog.List() {
		fmt.Printf("  - %s\n", id)
	}
}
