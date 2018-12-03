package input

import (
	"bufio"
	"os"
	"strconv"
)

type Reader struct {
	path string
	file os.File
	scanner bufio.Scanner
}

func NewReader(path string) Reader {
	file, err := os.Open(path)
	check(err)

	return Reader{
		path: path,
		file: *file,
		scanner: *bufio.NewScanner(file),
	}
}

func (r Reader) ReadLines() []string {
	if &r.file != nil {
		defer r.file.Close()
	}

	lines := []string{}

	for r.scanner.Scan() {
		lines = append(lines, r.scanner.Text())
	}

	return lines
}

func (r Reader) ReadLinesAsNumbers() []int {
	if &r.file != nil {
		defer r.file.Close()
	}

	numbers := []int{}

	for r.scanner.Scan() {
		converted, err := strconv.Atoi(r.scanner.Text())
		check(err)

		numbers = append(numbers, converted)
	}

	return numbers
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
