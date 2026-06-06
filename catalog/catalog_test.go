package catalog

import (
	"testing"
)

func TestList(t *testing.T) {
	catalogs := List()
	if len(catalogs) == 0 {
		t.Error("List() returned empty list")
	}

	// Check for expected catalogs
	expected := []string{
		"2.5-edition-wcag-2.2-508-en",
		"2.5-edition-wcag-2.1-508-en",
		"2.5-edition-wcag-2.0-508-en",
	}

	for _, id := range expected {
		found := false
		for _, c := range catalogs {
			if c == id {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected catalog %q not found in List()", id)
		}
	}
}

func TestGet(t *testing.T) {
	cat, err := Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if cat.Title == "" {
		t.Error("catalog title is empty")
	}
	if cat.Lang != "en" {
		t.Errorf("expected lang 'en', got '%s'", cat.Lang)
	}
	if len(cat.Chapters) == 0 {
		t.Error("catalog has no chapters")
	}
	if len(cat.Components) == 0 {
		t.Error("catalog has no components")
	}
	if len(cat.Standards) == 0 {
		t.Error("catalog has no standards")
	}
}

func TestGetCached(t *testing.T) {
	// First load
	cat1, err := Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Second load should return cached
	cat2, err := Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		t.Fatalf("Get() second call error = %v", err)
	}

	// Should be the same pointer
	if cat1 != cat2 {
		t.Error("Get() should return cached catalog")
	}
}

func TestGetNotFound(t *testing.T) {
	_, err := Get("nonexistent-catalog")
	if err == nil {
		t.Error("expected error for nonexistent catalog")
	}
}

func TestMustGet(t *testing.T) {
	cat := MustGet("2.5-edition-wcag-2.2-508-en")
	if cat == nil {
		t.Error("MustGet() returned nil")
	}
}

func TestMustGetPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustGet() should panic for nonexistent catalog")
		}
	}()

	MustGet("nonexistent-catalog")
}

func TestExists(t *testing.T) {
	if !Exists("2.5-edition-wcag-2.2-508-en") {
		t.Error("Exists() should return true for existing catalog")
	}
	if Exists("nonexistent-catalog") {
		t.Error("Exists() should return false for nonexistent catalog")
	}
}

func TestCatalogStructure(t *testing.T) {
	cat, err := Get("2.5-edition-wcag-2.2-508-en")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	// Check WCAG criteria exist
	var levelAChapter *ChapterDef
	for i := range cat.Chapters {
		if cat.Chapters[i].ID == "success_criteria_level_a" {
			levelAChapter = &cat.Chapters[i]
			break
		}
	}

	if levelAChapter == nil {
		t.Fatal("success_criteria_level_a chapter not found")
	}

	// Check for well-known criteria
	criterion111Found := false
	for _, c := range levelAChapter.Criteria {
		if c.ID == "1.1.1" {
			criterion111Found = true
			if c.Handle == "" {
				t.Error("criterion 1.1.1 has no handle")
			}
			if len(c.Components) == 0 {
				t.Error("criterion 1.1.1 has no components")
			}
			break
		}
	}

	if !criterion111Found {
		t.Error("criterion 1.1.1 not found in level A chapter")
	}
}

func TestAllCatalogsLoadable(t *testing.T) {
	for _, id := range List() {
		t.Run(id, func(t *testing.T) {
			cat, err := Get(id)
			if err != nil {
				t.Errorf("failed to load catalog %s: %v", id, err)
				return
			}

			if cat.Title == "" {
				t.Error("catalog has empty title")
			}
			if len(cat.Chapters) == 0 {
				t.Error("catalog has no chapters")
			}
		})
	}
}
