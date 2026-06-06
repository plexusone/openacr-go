package openacr

// ComponentName represents the type of product component being evaluated.
type ComponentName string

// Valid component names as defined by OpenACR.
const (
	ComponentWeb           ComponentName = "web"
	ComponentElectronicDoc ComponentName = "electronic-docs"
	ComponentSoftware      ComponentName = "software"
	ComponentAuthoringTool ComponentName = "authoring-tool"
	ComponentNone          ComponentName = "none"
)

// ValidComponentNames returns all valid component names.
func ValidComponentNames() []ComponentName {
	return []ComponentName{
		ComponentWeb,
		ComponentElectronicDoc,
		ComponentSoftware,
		ComponentAuthoringTool,
		ComponentNone,
	}
}

// IsValid returns true if the component name is valid.
func (c ComponentName) IsValid() bool {
	switch c {
	case ComponentWeb, ComponentElectronicDoc, ComponentSoftware, ComponentAuthoringTool, ComponentNone:
		return true
	default:
		return false
	}
}

// AdherenceLevel represents the level of accessibility conformance.
type AdherenceLevel string

// Valid adherence levels as defined by OpenACR/VPAT.
const (
	LevelSupports          AdherenceLevel = "supports"
	LevelPartiallySupports AdherenceLevel = "partially-supports"
	LevelDoesNotSupport    AdherenceLevel = "does-not-support"
	LevelNotApplicable     AdherenceLevel = "not-applicable"
	LevelNotEvaluated      AdherenceLevel = "not-evaluated"
)

// ValidAdherenceLevels returns all valid adherence levels.
func ValidAdherenceLevels() []AdherenceLevel {
	return []AdherenceLevel{
		LevelSupports,
		LevelPartiallySupports,
		LevelDoesNotSupport,
		LevelNotApplicable,
		LevelNotEvaluated,
	}
}

// IsValid returns true if the adherence level is valid.
func (a AdherenceLevel) IsValid() bool {
	switch a {
	case LevelSupports, LevelPartiallySupports, LevelDoesNotSupport, LevelNotApplicable, LevelNotEvaluated:
		return true
	default:
		return false
	}
}

// Component represents an evaluation of a specific component type.
type Component struct {
	// Name is the component type being evaluated.
	Name ComponentName `json:"name" yaml:"name"`

	// Adherence contains the conformance level and notes.
	Adherence Adherence `json:"adherence,omitempty" yaml:"adherence,omitempty"`
}

// Adherence represents the conformance status of a criterion for a component.
type Adherence struct {
	// Level is the conformance level.
	Level AdherenceLevel `json:"level" yaml:"level"`

	// Notes contains additional information about the conformance status.
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`
}
