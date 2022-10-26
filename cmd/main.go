package main

import (
	"flag"
	"fmt"
	"github.com/rnsloan/changelog/pkg/changelog"
)

/* TODO
- readme
- add large flag: git gc && git count-objects -vH (> 200000 git objects)
*/

func main() {
	repositoryPath := flag.String("repositoryPath", "", "local file path or remote URL to git repository. Defaults to current directory")
	outputPath := flag.String("outputPath", "", "output location for CHANGELOG.md. Defaults to current directory")
	formatMessage := flag.Bool("formatMessage", true, "start all commit message lines with a hyphen character. Default true")
	flag.Parse()
	fmt.Println(*repositoryPath)
	changelog.Build(&changelog.Config{RepositoryPath: *repositoryPath, OutputPath: *outputPath, FormatMessage: *formatMessage})
}
