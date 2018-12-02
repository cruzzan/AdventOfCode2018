package router

import "github.com/cruzzan/AdventOfCode2018/internal/pkg/days"

type Router interface {
	Route(day string, puzzle int)
}

type router struct {}

func NewRouter() Router {
	return &router{}
}

func (r router) Route(dayName string, puzzle int)  {
	dayFactory := days.Factory{}

	day := dayFactory.GetDay(dayName)

	day.Run(puzzle)
}
