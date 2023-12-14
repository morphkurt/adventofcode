package main

import (
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var m = make(map[string]string)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	r := 0
	northRotated := rotateNorth(input)
	for i, v := range strings.Split(northRotated, "\n") {
		r += (len(v) - i) * countRocks(v)
	}
	return r
}

func task2(input string) int {
	r := 0
	pattern := make(map[string]int)
	patternIter := make(map[int]string)
	cycleStart, cycleLength := 0, 0
	final := ""
	temp := input
	n := 1000000000
	for i := 1; i < n+1; i++ {
		if val, ok := m[temp]; ok {
			temp = val
		} else {
			rotateNorth := rotateNorth(temp)
			rotateWest := rotateWest(rotateNorth)
			rotateSouth := rotateSouth(rotateWest)
			final = rotateEast(rotateSouth)
			m[temp] = final
			temp = final
		}
		if _, ok := pattern[temp]; ok {
			cycleLength = i - pattern[temp]
			cycleStart = pattern[temp]
			break
		} else {
			pattern[temp] = i
			patternIter[i] = temp
		}
	}
	patternIndex := (n-cycleStart)%cycleLength + cycleStart
	temp = patternIter[patternIndex]
	for i, v := range strings.Split(temp, "\n") {
		r += (len(v) - i) * countRocks(v)
	}
	return r
}

func Parse(input string) (mirrors []string) {
	return strings.Split(input, "\n\n")
}

func transpose(frame string) string {
	a := [][]string{}
	for _, v := range strings.Split(frame, "\n") {
		a = append(a, strings.Split(v, ""))
	}
	o := util.Transpose(a)
	re := []string{}
	for _, v := range o {
		re = append(re, strings.Join(v, ""))
	}
	return strings.Join(re, "\n")
}

func rotateNorth(mirror string) string {
	//transpose (convert row > col)
	rotated := transpose(mirror)
	//move rocks to left
	v := rotate(rotated)
	//transpose (convert row > col)
	return transpose(v)
}

func rotateWest(mirror string) string {
	//move rocks to left
	v := rotate(mirror)
	return (v)
}

func rotateEast(mirror string) string {
	//move rocks to right
	v := rotateReverse(mirror)
	return (v)
}

func rotateSouth(mirror string) string {
	//transpose (convert row > col)
	rotated := transpose(mirror)
	//move rocks to right
	v := rotateReverse(rotated)
	//transpose (convert row > col)
	return transpose(v)
}

func rotate(mirror string) string {
	lines := strings.Split(mirror, "\n")
	out := []string{}
	for _, line := range lines {
		sections := strings.Split(line, "#")
		for k, s := range sections {
			if len(s) > 0 {
				newSection := ""
				rocks := countRocks(s)
				for i := len(s) - 1; i >= 0; i-- {
					if rocks > 0 {
						newSection += "O"
						rocks--
					} else {
						newSection += "."
					}
				}
				sections[k] = newSection
			}
		}
		out = append(out, strings.Join(sections, "#"))
	}
	return strings.Join(out, "\n")
}

func rotateReverse(mirror string) string {
	lines := strings.Split(mirror, "\n")
	l1 := []string{}
	for _, l := range lines {
		sections := strings.Split(l, "#")
		for k, s := range sections {
			if len(s) > 0 {
				movedSection := ""
				rocks := countRocks(s)
				for i := len(s) - 1; i >= 0; i-- {
					if rocks > 0 {
						movedSection = "O" + movedSection
						rocks--
					} else {
						movedSection = "." + movedSection
					}
				}
				sections[k] = movedSection
			}
		}
		l1 = append(l1, strings.Join(sections, "#"))
	}
	return strings.Join(l1, "\n")
}

func countRocks(line string) int {
	c := 0
	for _, v := range line {
		if v == 'O' {
			c++
		}
	}
	return c
}
