package main

import (
	"fmt"
	"slices"
	"sort"
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
	c := Parse(input)
	elvesCarry := []int{}
	for _, e := range c {
		total := 0
		for _, c := range e {
			total += c
		}
		elvesCarry = append(elvesCarry, total)
	}
	sort.Ints(elvesCarry)
	slices.Reverse(elvesCarry)
	return elvesCarry[0] + elvesCarry[1] + elvesCarry[2]
}

func task1(input string) int {
	c := Parse(input)
	maxCalories := 0
	for _, e := range c {
		total := 0
		for _, c := range e {
			total += c
		}
		if total > maxCalories {
			maxCalories = total
		}
	}
	return maxCalories
}
func Parse(input string) [][]int {
	groups := strings.Split(input, "\n\n")
	out := [][]int{}
	for _, v := range groups {
		carry := []int{}
		food := strings.Split(v, "\n")
		for _, c := range food {
			i, _ := strconv.Atoi(c)
			carry = append(carry, i)
		}
		out = append(out, carry)
	}
	return out
}
