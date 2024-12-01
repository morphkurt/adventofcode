package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
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
	c, m := ParseTask2(input)
	sum := 0
	for _, left := range c {
		sum += m[left] * left
	}
	return sum
}

func task1(input string) int {
	c := ParseTask1(input)
	left := []int{}
	right := []int{}
	for i := 0; i < len(c); i++ {
		left = append(left, c[i][0])
		right = append(right, c[i][1])
	}
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i := 0; i < len(c); i++ {
		sum += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	return sum
}
func ParseTask1(input string) [][]int {
	lines := strings.Split(input, "\n")
	out := [][]int{}
	for _, v := range lines {
		values := strings.Split(v, "   ")
		line := []int{}
		for _, c := range values {
			i, _ := strconv.Atoi(c)
			line = append(line, i)
		}
		out = append(out, line)
	}
	return out
}

func ParseTask2(input string) ([]int, map[int]int) {
	lines := strings.Split(input, "\n")
	m := make(map[int]int)
	out := []int{}
	for _, v := range lines {
		values := strings.Split(v, "   ")
		left, _ := strconv.Atoi(values[0])
		right, _ := strconv.Atoi(values[1])
		_, ok := m[right]
		if ok {
			m[right] = m[right] + 1
		} else {
			m[right] = 1
		}
		out = append(out, left)
	}
	return out, m
}
