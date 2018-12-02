package main

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/Menu"
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/Router"
)

func main()  {
	r := Router.NewRouter()
	m := Menu.NewMenu(r)

	m.Parse()
	m.Execute()
}
