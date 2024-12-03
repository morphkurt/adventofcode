package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task2(input string) int {
	c := parse(input)
	coef := 1
	sum := 0
	for _, v := range c {
		if strings.Contains(v, "don't") {
			coef = 0
		} else if strings.Contains(v, "do") {
			coef = 1
		} else if strings.Contains(v, "mul") {
			var a int
			var b int
			fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
			sum += coef * a * b
		}
	}
	return sum
}

func task1(input string) int {
	c := parse(input)
	coef := 1
	sum := 0
	for _, v := range c {
		if strings.Contains(v, "don't") {
			coef = 1
		} else if strings.Contains(v, "do") {
			coef = 1
		} else if strings.Contains(v, "mul") {
			var a int
			var b int
			fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
			sum += coef * a * b
		}
	}
	return sum
}

func parse(input string) []string {
	var re = regexp.MustCompile(`(?m)mul\(\d+,\d+\)|don't|do`)
	return re.FindAllString(input, -1)
}
