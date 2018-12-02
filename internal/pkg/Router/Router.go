package Router

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/days/one"
)

type Router interface {
	Route(day string)
}

type router struct {}

func NewRouter() Router {
	return &router{}
}

func (r router) Route(day string)  {
	switch day {
	case "one":
		one.Run("")
	default:
		fmt.Println("Gonna run all tasks, or something.")
	}
}
