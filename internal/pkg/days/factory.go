package days

import "github.com/cruzzan/AdventOfCode2018/internal/pkg/input"

type Factory struct {}

func (f Factory) GetDay(name string) Day {
	switch name {
	case "one":

		return One{
			input.NewReader("inputs/one/1.txt"),
		}
		break
	}

	return One {
		input.NewReader("inputs/one/1.txt"),
	}
}
