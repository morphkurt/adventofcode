package main

import (
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var m = make(map[string]int)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	file := Parse(input)
	v := 0
	for _, line := range file {
		r := 100 * splitPoint(line)
		c := splitPoint(transpose(line))
		if c > 0 {
			v += c
		} else {
			v += r
		}
	}
	return v
}

func task2(input string) int {
	file := Parse(input)
	v := 0
	for _, line := range file {
		r := 100 * splitPointWithError(line)
		c := splitPointWithError(transpose(line))
		if c > 0 {
			v += c
		} else {
			v += r
		}
	}
	return v
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

func splitPoint(grid string) int {
	mirror := strings.Split(grid, "\n")
	for i := 0; i < len(mirror); i++ {
		curr := i
		next := i + 1
		reflectionFound := false
		for curr >= 0 && next < len(mirror) {
			if mirror[curr] == mirror[next] {
				reflectionFound = true
			} else {
				reflectionFound = false
				break
			}
			curr--
			next++
		}
		if reflectionFound {
			return i + 1
		}
	}
	return 0
}

func splitPointWithError(grid string) int {
	mirror := strings.Split(grid, "\n")
	for i := 0; i < len(mirror); i++ {
		curr := i
		next := i + 1
		j := 0
		reflectionFound := false
		for curr >= 0 && next < len(mirror) {
			for i := range mirror[curr] {
				if mirror[curr][i] != mirror[next][i] {
					j++
				}
			}
			if j == 1 {
				reflectionFound = true
			} else {
				reflectionFound = false
			}
			curr--
			next++
		}
		if reflectionFound {
			return i + 1
		}
	}
	return 0
}
