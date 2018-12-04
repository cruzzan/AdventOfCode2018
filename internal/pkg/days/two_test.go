package days

import "testing"

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
			t.Errorf("Expected doubles check in %s to return %v but got %v", c.in, c.want, got)
		}
	}
}
