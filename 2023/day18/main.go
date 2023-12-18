package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var NORTH = []int{0, 0, -1}
var SOUTH = []int{1, 0, 1}
var EAST = []int{3, 1, 0}
var WEST = []int{4, -1, 0}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	ins := Parse(input)
	matrix := [][]int{}
	//starting location
	s := []int{0, 0}
	matrix = append(matrix, s)
	maxX := math.MinInt64
	minX := math.MaxInt32
	maxY := math.MinInt64
	minY := math.MaxInt32
	for _, i := range ins {
		curr := matrix[len(matrix)-1]
		matrix = append(matrix, []int{curr[0] + (i[2] * i[0]), curr[1] + (i[3] * i[0])})
		maxX = int(math.Max(float64(maxX), float64(matrix[len(matrix)-1][0])))
		maxY = int(math.Max(float64(maxY), float64(matrix[len(matrix)-1][1])))
		minX = int(math.Min(float64(minX), float64(matrix[len(matrix)-1][0])))
		minY = int(math.Min(float64(minY), float64(matrix[len(matrix)-1][1])))
	}

	area := 0
	j := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		area += (matrix[j][0] + matrix[i][0]) * (matrix[i][1] - matrix[j][1])
		j = i
	}

	perimeter := 0
	j = len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		perimeter += int(math.Abs(float64((matrix[j][0] - matrix[i][0]))) + math.Abs(float64((matrix[i][1] - matrix[j][1]))))
		j = i
	}
	return area/2 + perimeter/2 + 1

}

func task2(input string) int {
	ins := Parse(input)
	matrix := [][]int{}
	//starting location
	s := []int{0, 0}
	matrix = append(matrix, s)
	maxX := math.MinInt64
	minX := math.MaxInt32
	maxY := math.MinInt64
	minY := math.MaxInt32
	for _, i := range ins {
		curr := matrix[len(matrix)-1]
		matrix = append(matrix, []int{curr[0] + (i[6] * i[4]), curr[1] + (i[7] * i[4])})
		maxX = int(math.Max(float64(maxX), float64(matrix[len(matrix)-1][0])))
		maxY = int(math.Max(float64(maxY), float64(matrix[len(matrix)-1][1])))
		minX = int(math.Min(float64(minX), float64(matrix[len(matrix)-1][0])))
		minY = int(math.Min(float64(minY), float64(matrix[len(matrix)-1][1])))
	}

	area := 0
	j := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		area += (matrix[j][0] + matrix[i][0]) * (matrix[i][1] - matrix[j][1])
		j = i
	}

	perimeter := 0
	j = len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		perimeter += int(math.Abs(float64((matrix[j][0] - matrix[i][0]))) + math.Abs(float64((matrix[i][1] - matrix[j][1]))))
		j = i
	}
	return area/2 + perimeter/2 + 1
}

func Parse(input string) (out [][]int) {
	out = [][]int{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		c := strings.Split(row, " ")
		direction := c[0]
		tempD := []int{}
		distance, _ := strconv.Atoi(c[1])

		hex := c[2]
		switch direction {
		case "R":
			tempD = EAST
		case "L":
			tempD = WEST
		case "U":
			tempD = NORTH
		case "D":
			tempD = SOUTH
		}
		hexD := []int{}
		hex = strings.ReplaceAll(hex, "(", "")
		hex = strings.ReplaceAll(hex, ")", "")
		hex = strings.ReplaceAll(hex, "#", "")
		hexValue := hex[:5]
		value, _ := strconv.ParseInt(hexValue, 16, 64)

		dir := int(hex[5] - '0')
		switch dir {
		case 0:
			hexD = EAST
		case 1:
			hexD = SOUTH
		case 2:
			hexD = WEST
		case 3:
			hexD = NORTH
		}
		out = append(out, []int{distance, tempD[0], tempD[1], tempD[2], int(value), hexD[0], hexD[1], hexD[2]})
	}
	return out
}
