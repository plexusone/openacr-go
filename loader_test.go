package openacr

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test loading a valid YAML file
	report, err := Load("testdata/valid.yaml")
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if report.Title != "Lorem Ipsum" {
		t.Errorf("expected title 'Lorem Ipsum', got '%s'", report.Title)
	}
	if report.Product.Name != "Lorem Ipsum" {
		t.Errorf("expected product name 'Lorem Ipsum', got '%s'", report.Product.Name)
	}
	if report.Product.Version != "1.1" {
		t.Errorf("expected product version '1.1', got '%s'", report.Product.Version)
	}
	if report.Author.Email != "cicero@example.com" {
		t.Errorf("expected author email 'cicero@example.com', got '%s'", report.Author.Email)
	}
}

func TestLoadDrupal(t *testing.T) {
	// Test loading a real-world example
	report, err := Load("testdata/drupal-9.yaml")
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if !strings.Contains(report.Product.Name, "Drupal") {
		t.Errorf("expected product name to contain 'Drupal', got '%s'", report.Product.Name)
	}

	// Check that chapters are loaded
	if len(report.Chapters) == 0 {
		t.Error("expected chapters to be loaded")
	}
}

func TestLoadInvalidFormat(t *testing.T) {
	// Create a temp file with invalid extension
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(tmpFile, []byte("test"), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := Load(tmpFile)
	if err == nil {
		t.Error("expected error for invalid file format")
	}
	if err != ErrInvalidFormat {
		t.Errorf("expected ErrInvalidFormat, got %v", err)
	}
}

func TestLoadYAML(t *testing.T) {
	yaml := `
title: Test Report
product:
  name: Test Product
  version: "1.0"
author:
  email: test@example.com
`
	report, err := LoadYAML(strings.NewReader(yaml))
	if err != nil {
		t.Fatalf("LoadYAML() error = %v", err)
	}

	if report.Title != "Test Report" {
		t.Errorf("expected title 'Test Report', got '%s'", report.Title)
	}
	if report.Product.Name != "Test Product" {
		t.Errorf("expected product name 'Test Product', got '%s'", report.Product.Name)
	}
}

func TestLoadJSON(t *testing.T) {
	json := `{
		"title": "Test Report",
		"product": {
			"name": "Test Product",
			"version": "1.0"
		},
		"author": {
			"email": "test@example.com"
		}
	}`
	report, err := LoadJSON(strings.NewReader(json))
	if err != nil {
		t.Fatalf("LoadJSON() error = %v", err)
	}

	if report.Title != "Test Report" {
		t.Errorf("expected title 'Test Report', got '%s'", report.Title)
	}
	if report.Product.Name != "Test Product" {
		t.Errorf("expected product name 'Test Product', got '%s'", report.Product.Name)
	}
}

func TestLoadYAMLBytes(t *testing.T) {
	yaml := []byte(`
title: Test Report
product:
  name: Test Product
author:
  email: test@example.com
`)
	report, err := LoadYAMLBytes(yaml)
	if err != nil {
		t.Fatalf("LoadYAMLBytes() error = %v", err)
	}

	if report.Title != "Test Report" {
		t.Errorf("expected title 'Test Report', got '%s'", report.Title)
	}
}

func TestLoadJSONBytes(t *testing.T) {
	json := []byte(`{
		"title": "Test Report",
		"product": {"name": "Test Product"},
		"author": {"email": "test@example.com"}
	}`)
	report, err := LoadJSONBytes(json)
	if err != nil {
		t.Fatalf("LoadJSONBytes() error = %v", err)
	}

	if report.Title != "Test Report" {
		t.Errorf("expected title 'Test Report', got '%s'", report.Title)
	}
}

func TestRoundTrip(t *testing.T) {
	// Load YAML
	original, err := Load("testdata/valid.yaml")
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	// Write to JSON
	var jsonBuf bytes.Buffer
	if err := original.WriteJSON(&jsonBuf); err != nil {
		t.Fatalf("WriteJSON() error = %v", err)
	}

	// Load from JSON
	reloaded, err := LoadJSON(&jsonBuf)
	if err != nil {
		t.Fatalf("LoadJSON() error = %v", err)
	}

	// Compare
	if original.Title != reloaded.Title {
		t.Errorf("title mismatch: %s != %s", original.Title, reloaded.Title)
	}
	if original.Product.Name != reloaded.Product.Name {
		t.Errorf("product name mismatch: %s != %s", original.Product.Name, reloaded.Product.Name)
	}
	if original.Author.Email != reloaded.Author.Email {
		t.Errorf("author email mismatch: %s != %s", original.Author.Email, reloaded.Author.Email)
	}
}
