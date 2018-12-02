package menu

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/pkg/router"
	"gopkg.in/alecthomas/kingpin.v2"
)

var(
	day	= kingpin.Arg("Day", "The day to run, ex. fifteen").Required().String()
)

func Parse()  {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("hello, you want to run %s? \n", *day)
}

func Execute()  {
	router.Route(*day)
}
