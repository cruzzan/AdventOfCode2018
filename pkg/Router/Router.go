package Router

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/pkg/days/one"
)

type Router struct {

}

func NewRouter() Router {
	return Router{}
}

func (r Router) Route(day string)  {
	switch day {
	case "one":
		one.Run("")
	default:
		fmt.Println("Gonna run all tasks, or something.")
	}
}
