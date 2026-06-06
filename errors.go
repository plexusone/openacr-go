package openacr

import "errors"

// Common errors returned by openacr functions.
var (
	// ErrInvalidFormat is returned when a file has an unsupported format.
	ErrInvalidFormat = errors.New("invalid format: must be .yaml, .yml, or .json")

	// ErrEmptyProductName is returned when the product name is empty.
	ErrEmptyProductName = errors.New("product name is required")

	// ErrEmptyAuthorEmail is returned when the author email is empty.
	ErrEmptyAuthorEmail = errors.New("author email is required")

	// ErrInvalidCatalog is returned when a catalog ID is not found.
	ErrInvalidCatalog = errors.New("catalog not found")

	// ErrInvalidChapter is returned when a chapter ID is not in the catalog.
	ErrInvalidChapter = errors.New("chapter not found in catalog")

	// ErrInvalidCriterion is returned when a criterion is not in the catalog chapter.
	ErrInvalidCriterion = errors.New("criterion not found in catalog chapter")

	// ErrInvalidComponent is returned when a component name is not valid.
	ErrInvalidComponent = errors.New("invalid component name")

	// ErrInvalidAdherenceLevel is returned when an adherence level is not valid.
	ErrInvalidAdherenceLevel = errors.New("invalid adherence level")
)

// ValidationError represents a validation error with context.
type ValidationError struct {
	Field   string
	Message string
	Err     error
}

// Error implements the error interface.
func (e ValidationError) Error() string {
	if e.Field != "" {
		return e.Field + ": " + e.Message
	}
	return e.Message
}

// Unwrap returns the underlying error.
func (e ValidationError) Unwrap() error {
	return e.Err
}
