package search_test

import (
	"bytes"
	"testing"

	"github.com/qba73/search"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	bufIn := bytes.NewBufferString("hello\nworld\none\nmore\ntime\nworld")
	bufOut := &bytes.Buffer{}

	s, err := search.New(
		search.WithInput(bufIn),
		search.WithOutput(bufOut),
	)
	if err != nil {
		t.Fatal(err)
	}

	s.Search("world")

	got := bufOut.String()
	want := "world\nworld\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
