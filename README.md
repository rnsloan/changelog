# Changelog

Generate a `CHANGELOG.md` file from the git history of a respository.

## Usage

Download the appropriate executable from the releases page [https://github.com/rnsloan/changelog/releases](https://github.com/rnsloan/changelog/releases).

To run: `./changelog`.

### Options

`--repositoryPath` the local file path or remote URL to git repository. Default: the current directory

`--outputPath` output location for CHANGELOG.md. Default: the current directory

`--formatMessage` start all commit message lines with a hyphen character. Default: `true`

`--large` clone the git repository temporarily to disk instead of in-memory. Recommended for repositories with a large history (> 200000 git objects). Default: `false`

## Development

- `go run ./cmd/main.go`
- `go test ./...`

To build a new release:

1. add a new git tag
2. push the tag to GitHub
3. `make release`
4. create the new release in GitHub and add the files in `/build` generated from the previous step as assets
