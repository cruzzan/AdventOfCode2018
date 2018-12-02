package days

import "fmt"

type One struct {

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
}

func (d One) Puzzle2() {
	fmt.Println("Task 1")
}
