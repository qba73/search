package search

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type option func(*search) error

// WithInput sets input source for the searcher.
// It errors when the provided input is nil.
func WithInput(r io.Reader) option {
	return func(s *search) error {
		if r == nil {
			return errors.New("nil input reader")
		}
		s.input = r
		return nil
	}
}

// WithOutput sets output for the searcher.
// It errors when the provided output is nil.
func WithOutput(w io.Writer) option {
	return func(s *search) error {
		if w == nil {
			return errors.New("nil output reader")
		}
		s.output = w
		return nil
	}
}

type search struct {
	input  io.Reader
	output io.Writer
}

// New creates a searcher with optional input and output.
// It uses stdout and stdin as default values for input
// and output. It returns an error if one of the provided
// options errors.
func New(opts ...option) (search, error) {
	s := search{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, o := range opts {
		if err := o(&s); err != nil {
			return search{}, err
		}
	}
	return s, nil

}

// Search takes a string and outputs all lines of the input
// that contains the given string.
func (s search) Search(str string) {
	scaner := bufio.NewScanner(s.input)
	for scaner.Scan() {
		l := scaner.Text()
		if !strings.Contains(l, str) {
			continue
		}
		s.output.Write([]byte(fmt.Sprintf("%s\n", l)))
	}
}

// Search uses defult searcher for searching lines that contain the given string.
// Default searcher uses stdin and stdout for input and output.
// Search panics if the searcher with default settings returns an error.
func Search(term string) {
	s, err := New()
	if err != nil {
		panic("internal error")
	}
	s.Search(term)
}
