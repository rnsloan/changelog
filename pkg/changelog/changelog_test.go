package changelog

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage"
	"testing"
)

/*
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
func TestPerformance(t *testing.T) {
	cloneRepository(&Config{RepositoryPath: "https://github.com/torvalds/linux.git", Large: true})
	fmt.Println("COMPLETE")
}
*/

const PC = "PlainClone"
const C = "Clone"

var funcName string

type MockClone struct{}

func (c MockClone) PlainClone(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
	funcName = PC
	return nil, nil
}
func (c MockClone) Clone(s storage.Storer, worktree billy.Filesystem, o *git.CloneOptions) (*git.Repository, error) {
	funcName = C
	return nil, nil
}

func TestCloneRepository(t *testing.T) {
	c := MockClone{}

	cloneRepository(&Config{}, c)

	if funcName != C {
		t.Errorf("wanted function called to be %s, Got: %s", C, funcName)
	}

	cloneRepository(&Config{Large: true}, c)

	if funcName != PC {
		t.Errorf("wanted function called to be %s, Got: %s", PC, funcName)
	}
}

func TestCreateFile(t *testing.T) {
	want := "foo/" + FileName
	got := createFile("foo")

	if got != want {
		t.Errorf("want file path to be %s Got: %s", want, got)
	}
}

func TestFormatDate(t *testing.T) {
	commit := `commit da39a3ee5e6b4b0d3255bfef95601890afd80709
Author: XXXX <xxxx@gmail.com>
Date:   Wed Oct 26 03:14:25 2022 +1100

    feat: first build


CHANGELOG.md created
`
	want := "2022-10-26"
	got := formatDate(commit)

	if got != want {
		t.Errorf("want date to be %s Got: %s", want, got)
	}
}

func TestFormatMessage(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"* hello", "- hello"},
		{"hello", "- hello"},
		{"-hello", "- hello"},
	} {
		got := formatMessage(c.in)
		if got != c.want {
			t.Errorf("formatMessage(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestNewConfig(t *testing.T) {
	got := newConfig(&Config{FormatMessage: false})

	if got.FormatMessage {
		t.Errorf("want FormatMessage to be false Got: %t", got.FormatMessage)
	}
}
