package days

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/input"
	"strings"
)

type Two struct {
	r input.Reader
}

func (d Two) Run(puzzle int)  {
	if puzzle == Both {
		fmt.Println("Running full day")
		d.Puzzle1()
		d.Puzzle2()
	} else if puzzle == P1 {
		d.Puzzle1()
	} else if puzzle == P2 {
		d.Puzzle2()
	}
}

func (d Two) Puzzle1()  {
	fmt.Println("Task 1")
	lines := d.r.ReadLines()
	doubles, triples := 0, 0

	for _, in := range lines {
		if d.containsDoubles(in) {
			doubles++
		}
		if d.containsTriples(in) {
			triples++
		}
	}

	fmt.Printf("Resulting checksum (%d * %d): %d \n", doubles, triples, doubles * triples)
}

func (d Two) Puzzle2()  {
	fmt.Println("Todo")
}

func (d Two) containsDoubles(s string) bool {
	unique := strings.Map(func (r rune) rune {
		return r
	}, s)

	for _, char := range strings.Split(unique, "") {
		if strings.Count(s, char) == 2 {
			return true
		}
	}

	return false
}

func (d Two) containsTriples(s string) bool {
	unique := strings.Map(func (r rune) rune {
		return r
	}, s)

	for _, char := range strings.Split(unique, "") {
		if strings.Count(s, char) == 3 {
			return true
		}
	}

	return false
}
