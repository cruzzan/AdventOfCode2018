package main

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/menu"
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/router"
)

func main()  {
	r := router.NewRouter()
	m := menu.NewMenu(r)

	m.Parse()
	m.Execute()
}
