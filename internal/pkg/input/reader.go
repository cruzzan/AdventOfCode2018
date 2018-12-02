package input

import (
	"bufio"
	"os"
	"strconv"
)

type Reader struct {
	path string
}

func NewReader(path string) Reader {
	return Reader{
		path,
	}
}

func (r Reader) ReadLines() []string {
	lines := []string{}
	file, err := os.Open(r.path)
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (r Reader) LinesToInts(lines []string) []int {
	ints := []int{}

	for _, s := range lines {
		converted, err := strconv.Atoi(s)
		check(err)

		ints = append(ints, converted)
	}

	return ints
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}
