package days

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/input"
)

type Factory struct {}

func (f Factory) GetDay(name string) Day {
	switch name {
	case "one":
		return One{
			input.NewReader("inputs/one/1.txt"),
			&[]int{},
		}
		break
	case "two":
		return Two{
			input.NewReader("inputs/two/1.txt"),
		}
	}

	return One {
		input.NewReader("inputs/one/1.txt"),
		&[]int{},
	}
}
