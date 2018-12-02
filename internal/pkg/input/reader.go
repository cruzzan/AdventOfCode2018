package input

import (
	"bufio"
	"os"
	"strconv"
)

type Reader struct {
	path string
	lines []string
}

func NewReader(path string) Reader {
	return Reader{
		path,
		[]string{},
	}
}

func (r Reader) ReadLines() []string {
	file, err := os.Open(r.path)
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		r.lines = append(r.lines, scanner.Text())
	}

	return r.lines
}

func (r Reader) LinesToInts() []int {
	ints := []int{}

	for _, s := range r.lines {
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
