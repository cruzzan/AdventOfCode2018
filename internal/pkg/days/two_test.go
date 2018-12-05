package days

import (
	"sort"
	"testing"
)

func TestTwo_containsDoubles(t *testing.T) {
	cases := []struct{
		in string
		want bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", true},
		{"abcccd", false},
		{"aabcdd", true},
		{"abcdee", true},
		{"ababab", false},
		{"aaaaaa", false},
	}

	dayTwo := Two{}

	for _, c := range cases {
		got := dayTwo.containsDoubles(c.in)

		if got != c.want {
			t.Errorf("Expected doubles check in %s to return %v but got %v", c.in, c.want, got)
		}
	}
}

func TestTwo_containsTriples(t *testing.T) {
	cases := []struct{
		in string
		want bool
	}{
		{"abcdef", false},
		{"bababc", true},
		{"abbcde", false},
		{"abcccd", true},
		{"aabcdd", false},
		{"abcdee", false},
		{"ababab", true},
		{"aaaaaa", false},
	}

	dayTwo := Two{}

	for _, c := range cases {
		got := dayTwo.containsTriples(c.in)

		if got != c.want {
			t.Errorf("Expected triples check in %s to return %v but got %v", c.in, c.want, got)
		}
	}
}

func TestTwo_findAlmostMatching(t *testing.T) {
	cases := []struct{
		in []string
		want []string
	}{
		{[]string{"aaa", "bbb", "aba", "abc", "aga", "gaa"}, []string{"aaa", "aba", "aga", "abc", "gaa"}},
		{[]string{"aaa", "bbb", "ccc"}, []string{}},
	}

	dayTwo := Two{}

	for _, c := range cases {
		got := dayTwo.closeMatches(c.in)

		if len(got) != len(c.want) {
			t.Errorf("Expected close matches for %v but got %v",c.want, got)
		}

		var sortedWant sort.StringSlice
		sortedWant = c.want
		var sortedGot sort.StringSlice
		sortedGot = got

		sortedWant.Sort()
		sortedGot.Sort()
		for i, want := range sortedWant {
			if sortedGot[i] != want {
				t.Errorf("Expected close matches for %v but got %v",c.want, got)
			}
		}
	}
}

func TestTwo_overlappingChars(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"Hello", "There"}, ""},
		{[]string{"", "bbb"}, ""},
		{[]string{"bbb", ""}, ""},
		{[]string{"bbb", "bbb"}, "bbb"},
		{[]string{"bbb", "bcb"}, "bb"},
	}

	dayTwo := Two{}

	for _, c := range cases {
		got := dayTwo.overlappingChars(c.in[0], c.in[1])

		if got != c.want {
			t.Errorf("Failed asserting that %s matches %s", c.want, got)
		}
	}
}
