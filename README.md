# Changelog

Generate a `CHANGELOG.md` file from the git history of a repository ([https://keepachangelog.com/](https://keepachangelog.com/)).

## Example

```
# Changelog

## 2022-11-05
### Added
- all: fix comment typos

- Change-Id: Ic16824482142d4de4d0b949459e36505ee944ff7
- Reviewed-on: https://go-review.googlesource.com/c/go/+/448175
- Reviewed-by: Robert Griesemer <gri@google.com>
- Run-TryBot: Dan Kortschak <dan@kortschak.io>
- Auto-Submit: Robert Griesemer <gri@google.com>
- TryBot-Result: Gopher Robot <gobot@golang.org>
- Auto-Submit: Dan Kortschak <dan@kortschak.io>
- Auto-Submit: Ian Lance Taylor <iant@google.com>
- Run-TryBot: Ian Lance Taylor <iant@google.com>
- Reviewed-by: Ian Lance Taylor <iant@google.com>

## 2022-11-03
### Added
- reflect: rewrite value.Equal to avoid allocations

- For #46746 

- Change-Id: I75ddb9ce24cd3394186562dae156fef9fe2d55d3
- Reviewed-on: https://go-review.googlesource.com/c/go/+/447798
- Reviewed-by: Ian Lance Taylor <iant@google.com>
- Run-TryBot: Ian Lance Taylor <iant@google.com>
- Auto-Submit: Ian Lance Taylor <iant@google.com>
- Reviewed-by: Bryan Mills <bcmills@google.com>
- TryBot-Result: Gopher Robot <gobot@golang.org>
```

## Usage

Download the appropriate executable from the releases page [https://github.com/rnsloan/changelog/releases](https://github.com/rnsloan/changelog/releases).

To run: `./changelog`.

The 'Types of changes' heading is always `Added`. 

### Options

`--repositoryPath` the local file path or remote URL to git repository. Default: the current directory

`--outputPath` output location for CHANGELOG.md. Default: the current directory

`--formatMessage` start all commit message lines with a hyphen character. Default: `true`

`--large` clone the git repository temporarily to disk instead of in-memory. Recommended for repositories with a large history of over 200,000 git objects (`git gc && git count-objects -vH`). Default: `false`

## Development

- `go run ./cmd/main.go`
- `go test ./...`

To build a new release:

1. add a new git tag
2. push the tag to GitHub
3. `make release`
4. create the new release in GitHub and add the files in `/build` generated from the previous step as assets
