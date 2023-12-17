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

func travel(curX, curY int, dir []int, matrix [][]rune, tiles [][]int) {

	if curX < 0 || curX > len(matrix[0])-1 || curY > len(matrix)-1 || curY < 0 ||
		hasVisited(tiles[curY][curX], dir) {
		//Either reached corner or reached a loop

	} else {
		nextSteps := []nextStep{}
		tiles[curY][curX] = markTile(tiles[curY][curX], dir)
		switch c := matrix[curY][curX]; c {
		case '.':
			nextSteps = append(nextSteps, nextStep{curX + dir[1], curY + dir[2], dir})
		case '/':
			nextDir := []int{}
			switch d := dir[0]; d {
			case NORTH[0]:
				nextDir = EAST
			case SOUTH[0]:
				nextDir = WEST
			case EAST[0]:
				nextDir = NORTH
			case WEST[0]:
				nextDir = SOUTH
			}
			nextSteps = append(nextSteps, nextStep{curX + nextDir[1], curY + nextDir[2], nextDir})
		case '\\':
			nextDirection := []int{}
			switch d := dir[0]; d {
			case NORTH[0]:
				nextDirection = WEST
			case SOUTH[0]:
				nextDirection = EAST
			case EAST[0]:
				nextDirection = SOUTH
			case WEST[0]:
				nextDirection = NORTH
			}
			nextSteps = append(nextSteps, nextStep{curX + nextDirection[1], curY + nextDirection[2], nextDirection})
		case '-':
			nextDir1 := []int{}
			nextDir2 := []int{}
			switch d := dir[0]; d {
			case NORTH[0]:
				nextDir1 = WEST
				nextDir2 = EAST
				nextSteps = append(nextSteps, nextStep{curX + nextDir1[1], curY + nextDir1[2], nextDir1})
				nextSteps = append(nextSteps, nextStep{curX + nextDir2[1], curY + nextDir2[2], nextDir2})
			case SOUTH[0]:
				nextDir1 = WEST
				nextDir2 = EAST
				nextSteps = append(nextSteps, nextStep{curX + nextDir1[1], curY + nextDir1[2], nextDir1})
				nextSteps = append(nextSteps, nextStep{curX + nextDir2[1], curY + nextDir2[2], nextDir2})

			default:
				nextSteps = append(nextSteps, nextStep{curX + dir[1], curY + dir[2], dir})
			}
		case '|':
			nextDir1 := []int{}
			nextDir2 := []int{}
			switch d := dir[0]; d {
			case WEST[0]:
				nextDir1 = NORTH
				nextDir2 = SOUTH
				nextSteps = append(nextSteps, nextStep{curX + nextDir1[1], curY + nextDir1[2], nextDir1})
				nextSteps = append(nextSteps, nextStep{curX + nextDir2[1], curY + nextDir2[2], nextDir2})

			case EAST[0]:
				nextDir1 = NORTH
				nextDir2 = SOUTH
				nextSteps = append(nextSteps, nextStep{curX + nextDir1[1], curY + nextDir1[2], nextDir1})
				nextSteps = append(nextSteps, nextStep{curX + nextDir2[1], curY + nextDir2[2], nextDir2})

			default:
				nextSteps = append(nextSteps, nextStep{curX + dir[1], curY + dir[2], dir})
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
