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
			t.Errorf("The lenghts don't match, expected %d but got %d", len(c.want), len(got))
		}
		for i, val := range c.want {
			if val != got[i] {
				t.Errorf("The slices don't match, expected %v but got %v", c.want, got)
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
			t.Errorf("The sums don't match, expected %d but got %d", c.want, got)
		}
	}
}
