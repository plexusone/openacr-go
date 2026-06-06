package openacr

import (
	"github.com/plexusone/openacr-go/catalog"
)

// Validate performs basic validation on the report and returns any validation errors.
// This checks required fields and valid values without requiring a catalog.
func (r *Report) Validate() []ValidationError {
	var errs []ValidationError

	// Required: product name
	if r.Product.Name == "" {
		errs = append(errs, ValidationError{
			Field:   "product.name",
			Message: "product name is required",
			Err:     ErrEmptyProductName,
		})
	}

	// Required: author email
	if r.Author.Email == "" {
		errs = append(errs, ValidationError{
			Field:   "author.email",
			Message: "author email is required",
			Err:     ErrEmptyAuthorEmail,
		})
	}

	// Validate component names and adherence levels in chapters
	for chapterID, chapter := range r.Chapters {
		for i, criterion := range chapter.Criteria {
			for j, comp := range criterion.Components {
				if !comp.Name.IsValid() {
					errs = append(errs, ValidationError{
						Field:   fieldPath("chapters", chapterID, "criteria", i, "components", j, "name"),
						Message: "invalid component name: " + string(comp.Name),
						Err:     ErrInvalidComponent,
					})
				}
				if comp.Adherence.Level != "" && !comp.Adherence.Level.IsValid() {
					errs = append(errs, ValidationError{
						Field:   fieldPath("chapters", chapterID, "criteria", i, "components", j, "adherence", "level"),
						Message: "invalid adherence level: " + string(comp.Adherence.Level),
						Err:     ErrInvalidAdherenceLevel,
					})
				}
			}
		}
	}

	return errs
}

// ValidateAgainstCatalog validates the report against a specific catalog.
// This checks that all chapters and criteria exist in the catalog.
func (r *Report) ValidateAgainstCatalog(cat *catalog.Catalog) []ValidationError {
	// First run basic validation
	errs := r.Validate()

	// Build lookup maps from catalog
	chapterMap := make(map[string]*catalog.ChapterDef)
	for i := range cat.Chapters {
		ch := &cat.Chapters[i]
		chapterMap[ch.ID] = ch
	}

	// Validate chapters exist in catalog
	for chapterID, chapter := range r.Chapters {
		catChapter, ok := chapterMap[chapterID]
		if !ok {
			errs = append(errs, ValidationError{
				Field:   fieldPath("chapters", chapterID),
				Message: "chapter not found in catalog: " + chapterID,
				Err:     ErrInvalidChapter,
			})
			continue
		}

		// Build criterion lookup for this chapter
		criterionMap := make(map[string]bool)
		for _, crit := range catChapter.Criteria {
			criterionMap[crit.ID] = true
			if crit.AltID != "" {
				criterionMap[crit.AltID] = true
			}
		}

		// Validate criteria exist in chapter
		for i, criterion := range chapter.Criteria {
			if !criterionMap[criterion.Num] {
				errs = append(errs, ValidationError{
					Field:   fieldPath("chapters", chapterID, "criteria", i, "num"),
					Message: "criterion not found in catalog chapter: " + criterion.Num,
					Err:     ErrInvalidCriterion,
				})
			}
		}
	}

	return errs
}

// fieldPath builds a field path string for error messages.
func fieldPath(parts ...any) string {
	result := ""
	for i, part := range parts {
		switch v := part.(type) {
		case string:
			if i > 0 {
				result += "."
			}
			result += v
		case int:
			result += "[" + itoa(v) + "]"
		}
	}
	return result
}

// itoa converts an integer to a string without importing strconv.
func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	var digits []byte
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		digits = append([]byte{byte('0' + i%10)}, digits...)
		i /= 10
	}
	if neg {
		digits = append([]byte{'-'}, digits...)
	}
	return string(digits)
}
