package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type One struct {}

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
	input := readInputLines()
	freqChanges := d.stringsToInts(input)
	result := d.sumFreqChanges(freqChanges)

	fmt.Printf("Resulting frequency: %d \n", result)
}

func (d One) Puzzle2() {
	fmt.Println("Task 2")
}

func (d One) sumFreqChanges(changes []int) int {
	result := 0

	for _, change := range changes {
		result += change
	}

	return result
}

func (d One) stringsToInts (strings []string) []int {
	ints := []int{}

	for _, s := range strings {
		converted, err := strconv.Atoi(s)
		check(err)

		ints = append(ints, converted)
	}

	return ints
}

func readInputLines() []string {
	inputLines := []string{}
	file, err := os.Open("inputs/one/1.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}


	return inputLines
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}
