package openacr

// Chapter represents a chapter in an OpenACR report containing
// accessibility evaluation criteria.
type Chapter struct {
	// Notes contains chapter-level notes.
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`

	// Disabled indicates whether this chapter is disabled/not applicable.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Criteria is the list of criteria evaluations in this chapter.
	Criteria []Criterion `json:"criteria,omitempty" yaml:"criteria,omitempty"`
}

// Criterion represents an individual accessibility criterion evaluation.
type Criterion struct {
	// Num is the criterion number/identifier (e.g., "1.1.1").
	Num string `json:"num" yaml:"num"`

	// Components contains the evaluation results for each component type.
	Components []Component `json:"components,omitempty" yaml:"components,omitempty"`
}
