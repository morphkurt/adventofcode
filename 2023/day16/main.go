package main

import (
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var NORTH = []int{1, 0, -1}
var SOUTH = []int{2, 0, 1}
var EAST = []int{3, 1, 0}
var WEST = []int{4, -1, 0}

type nextStep struct {
	x         int
	y         int
	direction []int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	m, t := Parse(input)
	travel(0, 0, EAST, m, t)
	return countTiles(t)
}

func task2(input string) int {
	m, t := Parse(input)
	maxValue := 0
	for y := 0; y < len(m); y++ {
		travel(0, y, EAST, m, t)
		r := countTiles(t)
		if r > maxValue {
			maxValue = r

		}
		t = resetTiles(t)
		travel(len(m[0])-1, y, WEST, m, t)
		r = countTiles(t)
		if r > maxValue {
			maxValue = r

		}
		t = resetTiles(t)

	}
	for x := 0; x < len(m); x++ {
		travel(x, 0, SOUTH, m, t)
		r := countTiles(t)
		if r > maxValue {
			maxValue = r

		}
		t = resetTiles(t)
		travel(x, len(m)-1, NORTH, m, t)
		r = countTiles(t)
		if r > maxValue {
			maxValue = r

		}
		t = resetTiles(t)

	}
	return maxValue
}

func Parse(input string) (matrix [][]rune, tiles [][]int) {
	matrix = [][]rune{}
	tiles = [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matrix = append(matrix, []rune(line))
	}
	for y := 0; y < len(matrix); y++ {
		col := []int{}
		for x := 0; x < len(matrix[0]); x++ {
			col = append(col, 0)
		}
		tiles = append(tiles, col)
	}
	return matrix, tiles
}

func travel(currentX, currentY int, direction []int, matrix [][]rune, tiles [][]int) {

	if currentX < 0 || currentX > len(matrix[0])-1 || currentY > len(matrix)-1 || currentY < 0 ||
		hasVisited(tiles[currentY][currentX], direction) {
		//Either reached corner or reached a loop

	} else {
		nextSteps := []nextStep{}
		tiles[currentY][currentX] = markTile(tiles[currentY][currentX], direction)
		switch c := matrix[currentY][currentX]; c {
		case '.':
			nextSteps = append(nextSteps, nextStep{currentX + direction[1], currentY + direction[2], direction})
		case '/':
			nextDirection := []int{}
			switch d := direction[0]; d {
			case NORTH[0]:
				nextDirection = EAST
			case SOUTH[0]:
				nextDirection = WEST
			case EAST[0]:
				nextDirection = NORTH
			case WEST[0]:
				nextDirection = SOUTH
			}
			nextSteps = append(nextSteps, nextStep{currentX + nextDirection[1], currentY + nextDirection[2], nextDirection})
		case '\\':
			nextDirection := []int{}
			switch d := direction[0]; d {
			case NORTH[0]:
				nextDirection = WEST
			case SOUTH[0]:
				nextDirection = EAST
			case EAST[0]:
				nextDirection = SOUTH
			case WEST[0]:
				nextDirection = NORTH
			}
			nextSteps = append(nextSteps, nextStep{currentX + nextDirection[1], currentY + nextDirection[2], nextDirection})
		case '-':
			nextDirection1 := []int{}
			nextDirection2 := []int{}
			switch d := direction[0]; d {
			case NORTH[0]:
				nextDirection1 = WEST
				nextDirection2 = EAST
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection1[1], currentY + nextDirection1[2], nextDirection1})
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection2[1], currentY + nextDirection2[2], nextDirection2})
			case SOUTH[0]:
				nextDirection1 = WEST
				nextDirection2 = EAST
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection1[1], currentY + nextDirection1[2], nextDirection1})
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection2[1], currentY + nextDirection2[2], nextDirection2})

			default:
				nextSteps = append(nextSteps, nextStep{currentX + direction[1], currentY + direction[2], direction})
			}
		case '|':
			nextDirection1 := []int{}
			nextDirection2 := []int{}
			switch d := direction[0]; d {
			case WEST[0]:
				nextDirection1 = NORTH
				nextDirection2 = SOUTH
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection1[1], currentY + nextDirection1[2], nextDirection1})
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection2[1], currentY + nextDirection2[2], nextDirection2})

			case EAST[0]:
				nextDirection1 = NORTH
				nextDirection2 = SOUTH
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection1[1], currentY + nextDirection1[2], nextDirection1})
				nextSteps = append(nextSteps, nextStep{currentX + nextDirection2[1], currentY + nextDirection2[2], nextDirection2})

			default:
				nextSteps = append(nextSteps, nextStep{currentX + direction[1], currentY + direction[2], direction})
			}
		}
		for _, step := range nextSteps {
			travel(step.x, step.y, step.direction, matrix, tiles)
		}
	}
}

func markTile(in int, direction []int) int {
	return in | 1<<direction[0]
}

func hasVisited(in int, direction []int) bool {
	return (in >> direction[0] & 1) == 1
}

func printMatrix(matrix [][]rune) {
	for _, line := range matrix {
		fmt.Println(string(line))
	}
}

func countTiles(matrix [][]int) int {
	count := 0
	for _, line := range matrix {
		for _, c := range line {
			if c != 0 {
				count++
			}
		}
	}
	return count
}

func resetTiles(matrix [][]int) (o [][]int) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			matrix[y][x] = 0
		}
	}
	return matrix
}
