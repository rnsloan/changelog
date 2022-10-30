package changelog

import (
	"bufio"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

const FileName = "CHANGELOG.md"
const TempDirectoryName = "changelog-file-generation"

type Config struct {
	// local file path or remote URL to git repository. Defaults to current directory
	RepositoryPath string
	// output location for CHANGELOG.md. Defaults to current directory
	OutputPath string
	// start all commit message lines with a hyphen character. Default true
	FormatMessage bool
	// Clone the repository temporarily to disk instead of in-memory. Recommended for repositories where 'git gc && git count-objects -vH' is > 200,000 git objects
	Large bool
}

type CloneRepositories interface {
	PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error)
	Clone(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error)
}

type Clone struct{}

func (c Clone) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	return git.PlainClone(path, isBare, o)
}
func (c Clone) Clone(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
	return git.Clone(s, worktree, o)
}

func newConfig(userConfig *Config) Config {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	check(err)

	c := Config{RepositoryPath: path, OutputPath: "", FormatMessage: true, Large: false}

	if userConfig.RepositoryPath != "" {
		c.RepositoryPath = userConfig.RepositoryPath
	}

	if userConfig.OutputPath != "" {
		c.OutputPath = userConfig.OutputPath
	}

	if !userConfig.FormatMessage {
		c.FormatMessage = false
	}

	if userConfig.Large {
		c.Large = true
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
	regBulletPoint := regexp.MustCompile(`^([-*])\s?`)
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

func createFile(path string) string {
	if path != "" && !strings.HasSuffix(path, "/") {
		return path + "/" + FileName
	}
	return path + FileName

}

func cloneRepository(config *Config, clone CloneRepositories) (*git.Repository, string, error) {
	var r *git.Repository
	var err error
	dir := ""

	if config.Large {
		dir, err = os.MkdirTemp("", TempDirectoryName)

		if err != nil {
			return nil, "", err
		}

		r, err = clone.PlainClone(dir, false, &git.CloneOptions{
			URL:      config.RepositoryPath,
			Progress: os.Stdout,
		})
	} else {
		r, err = clone.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL:      config.RepositoryPath,
			Progress: os.Stdout,
		})
	}

	if err != nil && dir != "" {
		os.RemoveAll(dir)
	}

	return r, dir, err
}

func check(e error) {
	if e != nil {
		fmt.Println("Error occurred. Exiting...")
		panic(e)
	}
}

func Build(c *Config) {
	config := newConfig(c)

	fmt.Printf("Getting repository: %s...\n", config.RepositoryPath)
	r, dir, err := cloneRepository(&config, &Clone{})

	if dir != "" {
		defer func() {
			err = os.RemoveAll(dir)
		}()
	}

	check(err)

	fmt.Println("Creating CHANGELOG.md...")
	file := createFile(config.OutputPath)
	f, err := os.Create(file)
	check(err)

	defer f.Close()

	buffer := bufio.NewWriter(f)
	_, err = buffer.WriteString("# Changelog\n")
	check(err)

	var currentDate string
	cIter, err := r.Log(&git.LogOptions{})
	check(err)

	fmt.Println("Adding markdown...")
	err = cIter.ForEach(func(c *object.Commit) error {
		formattedDate := formatDate(c.String())
		markdown := c.Message

		if config.FormatMessage {
			markdown = formatMessage(c.Message)
		}

		if formattedDate != currentDate {
			// start a new markdown section on new date
			markdown = fmt.Sprintf("\n## %s\n### Added\n%s", formattedDate, markdown)
			currentDate = formattedDate
		}

		_, e := buffer.WriteString(markdown)
		if e != nil {
			return e
		}

		return nil
	})

	err = buffer.Flush()

	check(err)
	fmt.Printf("%s created\n", file)
}
