package menu

import (
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/router"
	"gopkg.in/alecthomas/kingpin.v2"
)

var(
	day	= kingpin.Arg("Day", "The day to run, ex. fifteen").Required().String()
)

type Menu struct {
	Router router.Router
} 

func NewMenu(r router.Router) Menu {
	return Menu{
		Router: r,
	}
}

func (m Menu) Parse()  {
	kingpin.Version("0.0.1")
	kingpin.Parse()
}

func (m Menu) Execute() {
	m.Router.Route(*day)
}
