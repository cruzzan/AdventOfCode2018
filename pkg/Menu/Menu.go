package Menu

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/pkg/Router"
	"gopkg.in/alecthomas/kingpin.v2"
)

var(
	day	= kingpin.Arg("Day", "The day to run, ex. fifteen").Required().String()
)

type Menu struct {
	Router Router.Router
} 

func NewMenu(r Router.Router) Menu {
	return Menu{
		Router: r,
	}
}

func (m Menu) Parse()  {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("hello, you want to run %s? \n", *day)
}

func (m Menu) Execute() {
	m.Router.Route(*day)
}
