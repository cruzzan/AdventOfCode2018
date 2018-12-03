package input

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

func TestReader_ReadLines(t *testing.T) {
	cases := []struct{
		in io.Reader
		want []string
	}{
		{strings.NewReader("-1\n-2\n9999"), []string{"-1", "-2", "9999"}},
		{strings.NewReader("111\n0\n-20583"), []string{"111", "0", "-20583"}},
		{strings.NewReader("0\n1\n100000000000"), []string{"0", "1", "100000000000"}},
		{strings.NewReader("Hello\nworld\n!"), []string{"Hello", "world", "!"}},
	}

	for _, c := range cases {
		r := Reader{
			scanner: *bufio.NewScanner(c.in),
		}

		got := r.ReadLines()

		// Check that the slices are the same length
		if len(got) != len(c.want) {
			t.Errorf("The expected input %v to yeild a slice of %d items, got %d items", c.in, len(c.want), len(got))
		}

		// Check that the slices contain the same elements, in order
		for i, val := range got {
			if val != c.want[i] {
				t.Errorf("Failed asserting that two slices are equal, expected %v got %v", c.want, got)
			}
		}
	}
}

func TestReader_ReadLinesAsNumbers(t *testing.T) {
	cases := []struct{
		in io.Reader
		want []int
	}{
		{strings.NewReader("-1\n-2\n9999"), []int{-1, -2, 9999}},
		{strings.NewReader("111\n0\n-20583"), []int{111, 0, -20583}},
		{strings.NewReader("0\n1\n100000000000"), []int{0, 1, 100000000000}},
	}

	for _, c := range cases {
		r := Reader{
			scanner: *bufio.NewScanner(c.in),
		}

		got := r.ReadLinesAsNumbers()

		// Check that the slices are the same length
		if len(got) != len(c.want) {
			t.Errorf("The expected input %v to yeild a slice of %d items, got %d items", c.in, len(c.want), len(got))
		}

		// Check that the slices contain the same elements, in order
		for i, val := range got {
			if val != c.want[i] {
				t.Errorf("Failed asserting that two slices are equal, expected %v got %v", c.want, got)
			}
		}
	}
}
