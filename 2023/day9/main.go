package main

import (
	"bufio"
	"fmt"
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

func task1(input string) int {
	rows := ParseInput(input)
	results := 0
	for _, row := range rows {
		intervals := [][]int{}
		intervals = append(intervals, row)
		tempRow := row
		for !CheckZeros(tempRow) {
			tempRow = GetSequence(tempRow)
			intervals = append(intervals, tempRow)
		}
		results += GetNextValue(intervals)
	}

	return results
}

func task2(input string) int {
	rows := ParseInput(input)
	results := 0
	for _, row := range rows {
		intervals := [][]int{}
		intervals = append(intervals, row)
		tempRow := row
		for !CheckZeros(tempRow) {
			tempRow = GetSequence(tempRow)
			intervals = append(intervals, tempRow)
		}
		results += GetPrevValue(intervals)
	}

	return results
}

func ParseInput(input string) (data [][]int) {
	result := [][]int{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		result = append(result, ToInt(strings.Split(scanner.Text(), " ")))
	}
	return result
}

func ToInt(in []string) (out []int) {
	for _, s := range in {
		i, _ := strconv.Atoi(s)
		out = append(out, i)
	}
	return out
}

func GetSequence(in []int) (out []int) {
	for i := 0; i < len(in)-1; i++ {
		curr := in[i]
		next := in[i+1]
		out = append(out, next-curr)
	}
	return out
}

func CheckZeros(in []int) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] != 0 {
			return false
		}
	}
	return true
}

func GetNextValue(in [][]int) int {
	lastValue := 0
	for i := len(in) - 1; i >= 0; i-- {
		lastValue += in[i][len(in[i])-1]
	}
	return lastValue
}

func GetPrevValue(in [][]int) int {
	lastValue := 0
	for i := len(in) - 1; i >= 0; i-- {
		lastValue = in[i][0] - lastValue
	}
	return lastValue
}
