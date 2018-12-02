package main

import (
	"github.com/cruzzan/AdventOfCode2018/pkg/Menu"
	"github.com/cruzzan/AdventOfCode2018/pkg/Router"
)

func main()  {
	r := Router.NewRouter()
	m := Menu.NewMenu(r)

	m.Parse()
	m.Execute()
}
