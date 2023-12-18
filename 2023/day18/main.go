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
	for _, i := range ins {
		curr := matrix[len(matrix)-1]
		matrix = append(matrix, []int{curr[0] + (i[2] * i[0]), curr[1] + (i[3] * i[0])})
	}
	return findArea(matrix)
}

func findArea(matrix [][]int) int {
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
	//til: pick's theorem
	return area/2 + perimeter/2 + 1
}

func task2(input string) int {
	ins := Parse(input)
	matrix := [][]int{}
	//starting location
	s := []int{0, 0}
	matrix = append(matrix, s)
	for _, i := range ins {
		curr := matrix[len(matrix)-1]
		matrix = append(matrix, []int{curr[0] + (i[6] * i[4]), curr[1] + (i[7] * i[4])})
	}
	return findArea(matrix)
}

func Parse(input string) (out [][]int) {
	out = [][]int{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		c := strings.Split(row, " ")
		direction := c[0]
		tempDirection := []int{}
		distance, _ := strconv.Atoi(c[1])

		hex := c[2]
		switch direction {
		case "R":
			tempDirection = EAST
		case "L":
			tempDirection = WEST
		case "U":
			tempDirection = NORTH
		case "D":
			tempDirection = SOUTH
		}
		hexDirections := []int{}
		hex = strings.ReplaceAll(hex, "(", "")
		hex = strings.ReplaceAll(hex, ")", "")
		hex = strings.ReplaceAll(hex, "#", "")
		hexValue := hex[:5]
		distanceFromHex, _ := strconv.ParseInt(hexValue, 16, 64)

		dir := int(hex[5] - '0')
		switch dir {
		case 0:
			hexDirections = EAST
		case 1:
			hexDirections = SOUTH
		case 2:
			hexDirections = WEST
		case 3:
			hexDirections = NORTH
		}
		out = append(out, []int{
			distance, // distance
			tempDirection[0],
			tempDirection[1],     // x
			tempDirection[2],     // y
			int(distanceFromHex), // distance from hex
			hexDirections[0],
			hexDirections[1],  //x from hex
			hexDirections[2]}) // y from hex
	}
	return out
}
