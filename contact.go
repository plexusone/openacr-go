package openacr

// Contact represents contact information for a person or organization.
type Contact struct {
	// Name is the name of the contact person.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// CompanyName is the name of the company or organization.
	CompanyName string `json:"company_name,omitempty" yaml:"company_name,omitempty"`

	// Address is the physical address.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Email is the email address.
	Email string `json:"email" yaml:"email"`

	// Phone is the phone number.
	Phone string `json:"phone,omitempty" yaml:"phone,omitempty"`

	// Website is the website URL.
	Website string `json:"website,omitempty" yaml:"website,omitempty"`
}
