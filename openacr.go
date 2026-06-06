package openacr

// Report represents an OpenACR accessibility conformance report.
type Report struct {
	// Title is the title of the report.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Product describes the product being evaluated.
	Product Product `json:"product" yaml:"product"`

	// Author is the person or organization that authored the report.
	Author Contact `json:"author" yaml:"author"`

	// Vendor is the product vendor (optional, may differ from author).
	Vendor *Contact `json:"vendor,omitempty" yaml:"vendor,omitempty"`

	// ReportDate is the date the report was created (YYYY-MM-DD format).
	ReportDate string `json:"report_date,omitempty" yaml:"report_date,omitempty"`

	// LastModifiedDate is the date the report was last modified (YYYY-MM-DD format).
	LastModifiedDate string `json:"last_modified_date,omitempty" yaml:"last_modified_date,omitempty"`

	// Version is the version number of the report.
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	// Notes contains general notes about the report.
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`

	// EvaluationMethods describes the evaluation methods used.
	EvaluationMethods string `json:"evaluation_methods_used,omitempty" yaml:"evaluation_methods_used,omitempty"`

	// LegalDisclaimer is the legal disclaimer for the report.
	LegalDisclaimer string `json:"legal_disclaimer,omitempty" yaml:"legal_disclaimer,omitempty"`

	// License is the license under which the report is published.
	License string `json:"license,omitempty" yaml:"license,omitempty"`

	// Repository is the URL of the repository containing the report.
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Feedback is information about how to provide feedback on the report.
	Feedback string `json:"feedback,omitempty" yaml:"feedback,omitempty"`

	// RelatedOpenACRs lists related OpenACR reports.
	RelatedOpenACRs []RelatedOpenACR `json:"related_openacrs,omitempty" yaml:"related_openacrs,omitempty"`

	// Catalog is the ID of the catalog used for this report.
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	// Chapters contains the accessibility evaluation results by chapter.
	Chapters map[string]Chapter `json:"chapters,omitempty" yaml:"chapters,omitempty"`
}

// Product describes the product being evaluated.
type Product struct {
	// Name is the name of the product.
	Name string `json:"name" yaml:"name"`

	// Version is the version of the product.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Description is a description of the product.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// RelatedOpenACR references another OpenACR report.
type RelatedOpenACR struct {
	// URL is the URL of the related report.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`

	// Type describes the relationship type.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// NewReport creates a new Report with the given options.
func NewReport(opts ...Option) *Report {
	r := &Report{
		Chapters: make(map[string]Chapter),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
