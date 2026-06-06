// Package catalog provides access to embedded OpenACR accessibility standard catalogs.
//
// Catalogs define the structure of accessibility standards including chapters,
// criteria, and valid components. The catalogs are embedded from the GSA OpenACR
// project and can be accessed by their ID.
//
// # Usage
//
//	cat, err := catalog.Get("2.5-edition-wcag-2.2-508-en")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Catalog: %s\n", cat.Title)
//
// # Available Catalogs
//
// Use List() to get all available catalog IDs. Common catalogs include:
//   - 2.5-edition-wcag-2.2-508-en: WCAG 2.2 with Section 508 (latest)
//   - 2.5-edition-wcag-2.1-508-en: WCAG 2.1 with Section 508
//   - 2.5-edition-wcag-2.0-508-en: WCAG 2.0 with Section 508
package catalog

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ErrCatalogNotFound is returned when a catalog ID is not found.
var ErrCatalogNotFound = errors.New("catalog not found")

// Catalog represents an OpenACR accessibility standard catalog.
type Catalog struct {
	// Title is the human-readable title of the catalog.
	Title string `json:"title" yaml:"title"`

	// Lang is the language code of the catalog (e.g., "en").
	Lang string `json:"lang" yaml:"lang"`

	// Standards lists the accessibility standards included in this catalog.
	Standards []Standard `json:"standards" yaml:"standards"`

	// Chapters contains the chapter definitions with their criteria.
	Chapters []ChapterDef `json:"chapters" yaml:"chapters"`

	// Components lists the valid component types for this catalog.
	Components []ComponentDef `json:"components" yaml:"components"`

	// Terms contains terminology definitions used in the catalog.
	Terms []Term `json:"terms" yaml:"terms"`
}

// Standard represents an accessibility standard included in a catalog.
type Standard struct {
	// ID is the unique identifier for the standard.
	ID string `json:"id" yaml:"id"`

	// Label is the human-readable name of the standard.
	Label string `json:"label" yaml:"label"`

	// URL is the link to the standard's documentation.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	// Chapters lists the chapter IDs that belong to this standard.
	Chapters []string `json:"chapters,omitempty" yaml:"chapters,omitempty"`
}

// ChapterDef represents a chapter definition in a catalog.
type ChapterDef struct {
	// ID is the unique identifier for the chapter.
	ID string `json:"id" yaml:"id"`

	// Label is the human-readable name of the chapter.
	Label string `json:"label" yaml:"label"`

	// Order is the display order of the chapter.
	Order int `json:"order" yaml:"order"`

	// Criteria contains the criterion definitions in this chapter.
	Criteria []CriterionDef `json:"criteria" yaml:"criteria"`
}

// CriterionDef represents a criterion definition in a catalog.
type CriterionDef struct {
	// ID is the unique identifier for the criterion (e.g., "1.1.1").
	ID string `json:"id" yaml:"id"`

	// Handle is the short name or handle for the criterion.
	Handle string `json:"handle" yaml:"handle"`

	// AltID is an alternative identifier for the criterion.
	AltID string `json:"alt_id,omitempty" yaml:"alt_id,omitempty"`

	// Components lists the component types applicable to this criterion.
	Components []string `json:"components" yaml:"components"`
}

// ComponentDef represents a component definition in a catalog.
type ComponentDef struct {
	// ID is the unique identifier for the component.
	ID string `json:"id" yaml:"id"`

	// Label is the human-readable name of the component.
	Label string `json:"label" yaml:"label"`
}

// Term represents a terminology definition in a catalog.
type Term struct {
	// ID is the unique identifier for the term.
	ID string `json:"id" yaml:"id"`

	// Label is the term being defined.
	Label string `json:"label" yaml:"label"`

	// Definition is the definition of the term.
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`
}

// LoadFromFile loads a catalog from a YAML file.
func LoadFromFile(path string) (*Catalog, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening catalog file: %w", err)
	}
	defer func() { _ = f.Close() }()

	return LoadFromReader(f)
}

// LoadFromReader loads a catalog from a YAML reader.
func LoadFromReader(r io.Reader) (*Catalog, error) {
	var cat Catalog
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&cat); err != nil {
		return nil, fmt.Errorf("decoding catalog YAML: %w", err)
	}
	return &cat, nil
}

// LoadFromBytes loads a catalog from YAML bytes.
func LoadFromBytes(data []byte) (*Catalog, error) {
	var cat Catalog
	if err := yaml.Unmarshal(data, &cat); err != nil {
		return nil, fmt.Errorf("unmarshaling catalog YAML: %w", err)
	}
	return &cat, nil
}
