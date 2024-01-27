package generator

import (
	"context"

	"github.com/lucastomic/web-generator/web-generator/internal/pagedata"
)

type Generator interface {
	// Generate takes the pagedata, geneartes the files (HTML, HTMX, CSS, etc.) and stores them in a temp folder.
	// It returns the paths of the generated files.
	GenerateAndGetPaths(context.Context, pagedata.PageData) ([]string, error)
}
