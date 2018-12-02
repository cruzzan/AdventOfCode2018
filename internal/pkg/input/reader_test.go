package input

import "testing"

func TestStringToIntConversion(t *testing.T) {
	cases := []struct{
		in []string
		want []int
	}{
		{[]string{"0"}, []int{0}},
		{[]string{"-1"}, []int{-1}},
		{[]string{"1"}, []int{1}},
		{[]string{"0", "2", "-3", "99999999"}, []int{0, 2, -3, 99999999}},
	}

	r := Reader{}

	for _, c := range cases {
		got := r.LinesToInts(c.in)
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
