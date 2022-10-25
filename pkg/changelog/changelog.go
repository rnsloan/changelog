package changelog

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const FileName = "CHANGELOG.md"

type Config struct {
	// local file path or remote URL to git repository. Defaults to current directory
	RepositoryPath string
	// output location for CHANGELOG.md. Defaults to current directory
	OutputPath string
}

func newConfig(userConfig *Config) Config {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	check(err)

	c := Config{RepositoryPath: path, OutputPath: ""}

	if userConfig.RepositoryPath != "" {
		c.RepositoryPath = userConfig.RepositoryPath
	}

	if userConfig.OutputPath != "" {
		c.OutputPath = userConfig.OutputPath
	}
	return c
}

func formatDate(commit string) string {
	split := strings.Split(commit, "\n")
	regDate := regexp.MustCompile(`Date:\s+`)

	// get date from commit
	date := string(regDate.ReplaceAll([]byte(split[2]), []byte("")))

	dt, _ := time.Parse("Mon Jan 02 15:04:05 2006 -0700", date)
	return dt.Format("2006-01-02")
}

func formatMessage(input string) string {
	regBulletPoint := regexp.MustCompile(`^(-|\*)\s?`)
	lines := strings.Split(input, "\n")
	formattedLines := make([]string, len(lines))

	for i, line := range lines {
		if line != "" {
			// remove any beginning '-' or "*" and whitespace characters
			striped := string(regBulletPoint.ReplaceAll([]byte(line), []byte("")))

			// have a consistent "- " prepended
			formatted := fmt.Sprintf("- %s", striped)

			formattedLines[i] = formatted
		}
	}

	return strings.Join(formattedLines, "\n")
}

func createFilePath(path string) string {
	if path != "" && !strings.HasSuffix(path, "/") {
		return path + "/" + FileName
	}
	return path + FileName

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Build(c *Config) {
	config := newConfig(c)

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: config.RepositoryPath,
	})

	check(err)

	f, err := os.Create(createFilePath(config.OutputPath))
	check(err)

	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("# Changelog\n")

	var currentDate string
	cIter, err := r.Log(&git.LogOptions{})
	check(err)

	err = cIter.ForEach(func(c *object.Commit) error {
		formattedDate := formatDate(c.String())
		markdown := formatMessage(c.Message)

		if formattedDate != currentDate {
			// start a new markdown section on new date
			markdown = fmt.Sprintf("\n## %s\n### Added\n%s", formattedDate, markdown)
			currentDate = formattedDate
		}

		_, err2 := w.WriteString(markdown)
		if err2 != nil {
			return err2
		}

		return nil
	})

	w.Flush()

	check(err)
}
