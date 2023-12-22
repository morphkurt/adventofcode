package main

import (
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var SOUTH = []int{0, 1}
var NORTH = []int{0, -1}
var EAST = []int{1, 0}
var WEST = []int{-1, 0}

type Point struct {
	x int
	y int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input, 64)
	fmt.Printf("task1:%d\n", task1)
	// task2 := task2(input, 1000)
	// fmt.Printf("task2:%d\n", task2)
}

func task1(input string, steps int) int {
	m := Parse(input)
	return solve(m, steps)
}

func solve(m [][]rune, steps int) int {
	count := 0
	s := FindStartingPoints(m)
	origin := s
	if len(s) > 0 {
		m[s[0].y][s[0].x] = '.'
	}
	for i := 1; i <= steps; i++ {
		moves := []Point{}
		for _, e := range s {
			moves = append(moves, e.FindNextStepsWithInfiniteMap(m)...)
		}
		uniqueMoves := RemoveDuplicates(moves)
		s = uniqueMoves
		count = len(s)
	}
	//revert the starting tile (hack)
	m[origin[0].y][origin[0].x] = 'S'
	return count
}

func Parse(input string) [][]rune {
	out := [][]rune{}
	rows := strings.Split(input, "\n")
	for _, v := range rows {
		out = append(out, []rune(v))
	}
	return out
}

func FindStartingPoints(m [][]rune) []Point {
	out := []Point{}
	for y, row := range m {
		for x := range row {
			if m[y][x] == 'S' || m[y][x] == 'O' {
				out = append(out, Point{x, y})
			}
		}
	}
	return out
}

func MarkMove(s, n Point, m [][]rune) {
	m[s.y][s.x] = '.'
	m[n.y][n.x] = 'O'
}

func (p Point) FindNextSteps(m [][]rune) []Point {
	out := []Point{}
	s := Point{p.x + SOUTH[0], p.y + SOUTH[1]}
	n := Point{p.x + NORTH[0], p.y + NORTH[1]}
	e := Point{p.x + EAST[0], p.y + EAST[1]}
	w := Point{p.x + WEST[0], p.y + WEST[1]}
	c := []Point{s, n, e, w}
	for _, e := range c {

		if e.x >= 0 && e.x < len(m[0]) && e.y >= 0 && e.y < len(m) {
			if m[e.y][e.x] == '.' || m[e.y][e.x] == 'S' {
				out = append(out, e)
			}
		}
	}
	return out
}

func (p Point) FindNextStepsWithInfiniteMap(m [][]rune) []Point {
	out := []Point{}
	s := Point{p.x + SOUTH[0], p.y + SOUTH[1]}
	n := Point{p.x + NORTH[0], p.y + NORTH[1]}
	e := Point{p.x + EAST[0], p.y + EAST[1]}
	w := Point{p.x + WEST[0], p.y + WEST[1]}
	c := []Point{s, n, e, w}
	for i, e := range c {
		if e.x < 0 {
			e.x = (e.x%len(m[0]) + len(m[0]))
		}
		if e.y < 0 {
			e.y = (e.y%len(m) + len(m))
		}

		e.x = e.x % len(m[0])
		e.y = e.y % len(m)
		if m[e.y][e.x] == '.' || m[e.y][e.x] == 'S' {
			out = append(out, c[i])
		}
	}
	return out
}

func RemoveDuplicates(points []Point) []Point {
	if len(points) == 0 {
		return points
	}
	allKeys := make(map[Point]bool)
	list := []Point{}
	for _, item := range points {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Count(m [][]rune) int {
	c := 0
	for y, row := range m {
		for x := range row {
			if m[y][x] == 'O' {
				c++
			}
		}
	}
	return c
}

func PrintMap(m [][]rune) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}
