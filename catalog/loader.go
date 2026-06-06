package catalog

import (
	"fmt"
	"path"
	"strings"
)

// catalogs caches loaded catalogs by ID.
var catalogs = make(map[string]*Catalog)

// catalogIDs lists all available catalog IDs.
var catalogIDs []string

func init() {
	// Discover available catalogs from embedded filesystem
	entries, err := catalogFS.ReadDir("data")
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasSuffix(name, ".yaml") {
			id := strings.TrimSuffix(name, ".yaml")
			catalogIDs = append(catalogIDs, id)
		}
	}
}

// List returns the IDs of all available catalogs.
func List() []string {
	result := make([]string, len(catalogIDs))
	copy(result, catalogIDs)
	return result
}

// Get loads and returns a catalog by its ID.
// The catalog is cached after the first load.
func Get(id string) (*Catalog, error) {
	// Check cache first
	if cat, ok := catalogs[id]; ok {
		return cat, nil
	}

	// Load from embedded filesystem
	filename := id + ".yaml"
	data, err := catalogFS.ReadFile(path.Join("data", filename))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrCatalogNotFound, id)
	}

	cat, err := LoadFromBytes(data)
	if err != nil {
		return nil, fmt.Errorf("loading catalog %s: %w", id, err)
	}

	// Cache and return
	catalogs[id] = cat
	return cat, nil
}

// MustGet loads and returns a catalog by its ID, panicking on error.
// This is useful for initialization code where the catalog must exist.
func MustGet(id string) *Catalog {
	cat, err := Get(id)
	if err != nil {
		panic(err)
	}
	return cat
}

// Exists returns true if a catalog with the given ID exists.
func Exists(id string) bool {
	for _, cid := range catalogIDs {
		if cid == id {
			return true
		}
	}
	return false
}
