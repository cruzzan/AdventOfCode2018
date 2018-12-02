package days

const Both int = 0
const P1 int = 1
const P2 int = 2

type Day interface {
	Run(puzzle int)
	Puzzle1()
	Puzzle2()
}
