package days

import (
	"testing"
)

func TestConversion(t *testing.T) {
	cases := []struct{
		in []string
		want []int
	}{
		{[]string{"0"}, []int{0}},
		{[]string{"-1"}, []int{-1}},
		{[]string{"1"}, []int{1}},
		{[]string{"0", "2", "-3", "99999999"}, []int{0, 2, -3, 99999999}},
	}

	dayOne := One{}

	for _, c := range cases {
		got := dayOne.stringsToInts(c.in)
		if len(got) != len(c.want) {
			t.Errorf("The lenghts don't match, expected %d but got %d\n", len(c.want), len(got))
		}
		for i, val := range c.want {
			if val != got[i] {
				t.Errorf("The slices don't match, expected %v but got %v\n", c.want, got)
			}
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct{
		in []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{-1}, -1},
		{[]int{1}, 1},
		{[]int{0, 2, -3, 99999999}, 99999998},
	}

	dayOne := One{}

	for _, c := range cases {
		got := dayOne.sumFreqChanges(c.in)

		if got != c.want {
			t.Errorf("The sums don't match, expected %d but got %d\n", c.want, got)
		}
	}
}

func TestFindRepeating(t *testing.T) {
	cases := []struct{
		in []int
		want int
	}{
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	dayOne := One{}

	for _, c := range cases {
		got := dayOne.findFirstFrequencyRepetition(c.in)

		if got != c.want {
			t.Errorf("The sums don't match, expected %d but got %d\n", c.want, got)
		}
	}
}

func TestNumberInSlice(t *testing.T) {
	cases := []struct{
		haystack []int
		needle int
		want bool
	}{
		{[]int{0, 1}, 0, true},
		{[]int{0}, 10, false},
		{[]int{0}, 0, true},
		{[]int{0, 10, -14, -5, 6}, 6, true},
		{[]int{0, 10, -14, -5, 6}, 7, false},
	}

	dayOne := One{}

	for _, c := range cases {
		got := dayOne.numberInSlice(c.haystack, c.needle)

		if got != c.want {
			t.Errorf("Unexpected result, expected %d precence in %v, to be %v got %v\n", c.needle, c.haystack, c.want, got)
		}
	}
}
