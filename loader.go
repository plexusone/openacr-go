package openacr

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Load reads an OpenACR report from a file.
// The file format is determined by the file extension (.yaml, .yml, or .json).
func Load(path string) (*Report, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer func() { _ = f.Close() }()

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".yaml", ".yml":
		return LoadYAML(f)
	case ".json":
		return LoadJSON(f)
	default:
		return nil, ErrInvalidFormat
	}
}

// LoadYAML reads an OpenACR report from a YAML reader.
func LoadYAML(r io.Reader) (*Report, error) {
	var report Report
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&report); err != nil {
		return nil, fmt.Errorf("decoding YAML: %w", err)
	}
	return &report, nil
}

// LoadJSON reads an OpenACR report from a JSON reader.
func LoadJSON(r io.Reader) (*Report, error) {
	var report Report
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&report); err != nil {
		return nil, fmt.Errorf("decoding JSON: %w", err)
	}
	return &report, nil
}

// LoadYAMLBytes reads an OpenACR report from YAML bytes.
func LoadYAMLBytes(data []byte) (*Report, error) {
	var report Report
	if err := yaml.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("unmarshaling YAML: %w", err)
	}
	return &report, nil
}

// LoadJSONBytes reads an OpenACR report from JSON bytes.
func LoadJSONBytes(data []byte) (*Report, error) {
	var report Report
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("unmarshaling JSON: %w", err)
	}
	return &report, nil
}
