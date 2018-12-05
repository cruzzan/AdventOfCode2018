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
	fmt.Println("Task 2")
	lines := d.r.ReadLines()
	closeMatches := d.closeMatches(lines)
	idOverlap := d.overlappingChars(closeMatches[0], closeMatches[1])
	fmt.Printf("Overlapping letters in ids: %s \n", idOverlap)
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

func (d Two) closeMatches(boxIds []string) []string {
	closeMatches := make(map[string]string)

	for _, boxId := range boxIds {
		for _, temp := range boxIds {
			misses := 0
			rtemp := []rune(temp)
			for i, char := range []rune(boxId) {
				if char != rtemp[i] {
					misses++
				}
			}

			if misses == 1 {
				closeMatches[boxId] = boxId
			}
		}
	}

	var result []string

	for _, match := range closeMatches {
		result = append(result, match)
	}

	return result
}

func (d Two) overlappingChars(s1 string, s2 string) string {
	var overlap []rune
	for i, val := range []rune(s1){
		if i < len(s2) {
			char2 := rune(s2[i])
			if val == char2 {
				overlap = append(overlap, char2)
			}
		}
	}

	return string(overlap)
}
