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

// Save writes the report to a file.
// The file format is determined by the file extension (.yaml, .yml, or .json).
func (r *Report) Save(path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	var data []byte
	var err error

	switch ext {
	case ".yaml", ".yml":
		data, err = r.YAML()
	case ".json":
		data, err = r.JSON()
	default:
		return ErrInvalidFormat
	}

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// WriteYAML writes the report as YAML to the given writer.
func (r *Report) WriteYAML(w io.Writer) error {
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(2)
	if err := encoder.Encode(r); err != nil {
		return fmt.Errorf("encoding YAML: %w", err)
	}
	return encoder.Close()
}

// WriteJSON writes the report as JSON to the given writer.
func (r *Report) WriteJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(r); err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}
	return nil
}

// YAML returns the report as YAML bytes.
func (r *Report) YAML() ([]byte, error) {
	data, err := yaml.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("marshaling YAML: %w", err)
	}
	return data, nil
}

// JSON returns the report as JSON bytes.
func (r *Report) JSON() ([]byte, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshaling JSON: %w", err)
	}
	return data, nil
}
