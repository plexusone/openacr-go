package openacr

import "time"

// Option is a functional option for configuring a Report.
type Option func(*Report)

// WithTitle sets the report title.
func WithTitle(title string) Option {
	return func(r *Report) {
		r.Title = title
	}
}

// WithProduct sets the product name and version.
func WithProduct(name, version string) Option {
	return func(r *Report) {
		r.Product = Product{
			Name:    name,
			Version: version,
		}
	}
}

// WithProductDescription sets the product with name, version, and description.
func WithProductDescription(name, version, description string) Option {
	return func(r *Report) {
		r.Product = Product{
			Name:        name,
			Version:     version,
			Description: description,
		}
	}
}

// WithAuthor sets the author name and email.
func WithAuthor(name, email string) Option {
	return func(r *Report) {
		r.Author = Contact{
			Name:  name,
			Email: email,
		}
	}
}

// WithAuthorCompany sets the author with company information.
func WithAuthorCompany(name, email, companyName string) Option {
	return func(r *Report) {
		r.Author = Contact{
			Name:        name,
			Email:       email,
			CompanyName: companyName,
		}
	}
}

// WithVendor sets the vendor contact information.
func WithVendor(name, email string) Option {
	return func(r *Report) {
		r.Vendor = &Contact{
			Name:  name,
			Email: email,
		}
	}
}

// WithCatalog sets the catalog ID for the report.
func WithCatalog(catalogID string) Option {
	return func(r *Report) {
		r.Catalog = catalogID
	}
}

// WithReportDate sets the report date.
func WithReportDate(date string) Option {
	return func(r *Report) {
		r.ReportDate = date
	}
}

// WithReportDateNow sets the report date to today.
func WithReportDateNow() Option {
	return func(r *Report) {
		r.ReportDate = time.Now().Format("2006-01-02")
	}
}

// WithVersion sets the report version number.
func WithVersion(version int) Option {
	return func(r *Report) {
		r.Version = version
	}
}

// WithNotes sets the report notes.
func WithNotes(notes string) Option {
	return func(r *Report) {
		r.Notes = notes
	}
}

// WithEvaluationMethods sets the evaluation methods used.
func WithEvaluationMethods(methods string) Option {
	return func(r *Report) {
		r.EvaluationMethods = methods
	}
}

// WithLegalDisclaimer sets the legal disclaimer.
func WithLegalDisclaimer(disclaimer string) Option {
	return func(r *Report) {
		r.LegalDisclaimer = disclaimer
	}
}

// WithLicense sets the license.
func WithLicense(license string) Option {
	return func(r *Report) {
		r.License = license
	}
}

// WithRepository sets the repository URL.
func WithRepository(url string) Option {
	return func(r *Report) {
		r.Repository = url
	}
}

// WithFeedback sets the feedback information.
func WithFeedback(feedback string) Option {
	return func(r *Report) {
		r.Feedback = feedback
	}
}
