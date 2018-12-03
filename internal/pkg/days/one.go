package days

import (
	"fmt"
	"github.com/cruzzan/AdventOfCode2018/internal/pkg/input"
)

type One struct {
	r input.Reader
	lines *[]int
}

func (d One) Run(puzzle int)  {
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

func (d One) Puzzle1() {
	fmt.Println("Task 1")
	lines := d.readLines()
	result := d.sumFreqChanges(lines)

	fmt.Printf("Resulting frequency: %d \n", result)
}

func (d One) Puzzle2() {
	fmt.Println("Task 2")
	lines := d.readLines()
	result := d.findFirstFrequencyRepetition(lines)

	fmt.Printf("First repeating frequency: %d \n", result)
}

func (d One) readLines() []int {
	if len(*d.lines) == 0 {
		*d.lines = d.r.ReadLinesAsNumbers()
	}

	return *d.lines
}

func (d One) sumFreqChanges(changes []int) int {
	result := 0

	for _, change := range changes {
		result += change
	}

	return result
}

func (d One) findFirstFrequencyRepetition(changes []int) int {
	result := 0
	frequencies := []int{result}

	for i := 0; i < len(changes); i++ {
		result += changes[i]

		if d.numberInSlice(frequencies, result) {
			return result
		} else {
			frequencies = append(frequencies, result)
		}

		if i == len(changes) - 1 {
			i = -1
		}
	}
	return 0
}

func (d One) numberInSlice(haystack[]int, n int) bool {
	for _, val := range haystack {
		if val == n {
			return true
		}
	}

	return false
}
