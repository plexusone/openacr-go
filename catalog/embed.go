package catalog

import "embed"

//go:embed data/*.yaml
var catalogFS embed.FS
