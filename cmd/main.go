package main

import (
	"flag"
	"github.com/rnsloan/changelog/pkg/changelog"
)

/* TODO
- readme
*/

func main() {
	repositoryPath := flag.String("repositoryPath", "", "local file path or remote URL to git repository. Defaults to current directory")
	outputPath := flag.String("outputPath", "", "output location for CHANGELOG.md. Defaults to current directory")
	formatMessage := flag.Bool("formatMessage", true, "start all commit message lines with a hyphen character. Default true")
	large := flag.Bool("large", false, "Clone the git repository temporarily to disk instead of in-memory. Recommended for repositories with a large history (> 200000 git objects). Default false")
	flag.Parse()

	changelog.Build(&changelog.Config{RepositoryPath: *repositoryPath, OutputPath: *outputPath, FormatMessage: *formatMessage, Large: *large})
}
