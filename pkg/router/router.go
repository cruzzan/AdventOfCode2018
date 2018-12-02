package router

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/pkg/days/one"
)

func Route(day string)  {
	switch day {
	case "one":
		one.Run("")
	default:
		fmt.Println("Gonna run all tasks, or something.")
	}
}
