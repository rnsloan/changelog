package changelog

import "testing"

func TestCreateFilePath(t *testing.T) {
	want := "foo/" + FileName
	got := createFilePath("foo")

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
