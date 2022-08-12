// Package search provides functions and structs necessary to search an array of files and print
// information on version descriptors that match obsolete files.
package search

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/kennethdyer/versc/file"
)

// RunSearch performs the search operation called from the command-line.
// It retrieves a list of files from the first argument and a product name
// and list of acceptabl eversions from the []string, then prints its findings
// to stdout.
func RunSearch(target string, products []string) {
	files := file.GetFiles(target)
	rets := collectData(files, products)

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprint(w, strings.Join(rets, "\n"))
	w.Flush()
}
