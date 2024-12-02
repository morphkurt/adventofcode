package main

import (
	"fmt"
	"math"
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
	c := parse(input)
	sum := 0
	for _, v := range c {
		if is_safe(v) {
			sum += 1
		} else {
			for i := 0; i < len(v); i++ {
				if is_safe(remove(v, i)) {
					sum += 1
					break
				}
			}
		}
	}
	return sum
}

func task1(input string) int {
	c := parse(input)
	sum := 0
	for _, v := range c {
		if is_safe(v) {
			sum += 1
		}
	}
	return sum
}

func is_safe(v []int) bool {
	safe := true
	for i := 2; i < len(v); i++ {
		prev_diff := v[1] - v[0]
		curr_diff := v[i] - v[i-1]
		if prev_diff*curr_diff <= 0 || math.Abs(float64(prev_diff)) > 3 || math.Abs(float64(curr_diff)) > 3 {
			safe = false
			break
		}
		safe = safe && true
	}
	return safe
}

func parse(input string) [][]int {
	lines := strings.Split(input, "\n")
	out := [][]int{}
	for _, v := range lines {
		if len(v) > 0 {
			values := strings.Split(v, " ")
			line := []int{}
			for _, c := range values {
				i, _ := strconv.Atoi(c)
				line = append(line, i)
			}
			out = append(out, line)
		}
	}
	return out
}

func remove(slice []int, index int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
}
