package main

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var SOUTH = []int{0, 1}
var NORTH = []int{0, -1}
var EAST = []int{1, 0}
var WEST = []int{-1, 0}

var Reset = "\033[0m"
var Red = "\033[31m"

var DIRECTIONS = [][]int{SOUTH, NORTH, EAST, WEST}

type Point struct {
	x int
	y int
}

type Node struct {
	p     Point
	steps int
}

type Nodes []Node

func (h Nodes) Len() int { return len(h) }
func (h Nodes) Less(i, j int) bool {

	return h[i].steps < h[j].steps
}
func (h Nodes) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Nodes) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *Nodes) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	m := parse(input)
	start := find('S', m)
	m[start.y][start.x] = 'a'
	end := find('E', m)
	m[end.y][end.x] = 'z'
	return solvePart1(m, start, end)
}

func task2(input string) int {
	m := parse(input)
	start := find('S', m)
	m[start.y][start.x] = 'a'
	end := find('E', m)
	m[end.y][end.x] = 'z'
	return solvePart2(m, end)
}

func Contains(c []Point, e Point) bool {
	for _, p := range c {
		if p.x == e.x && p.y == e.y {
			return true
		}
	}
	return false
}

func solvePart1(m [][]rune, start, end Point) int {
	h := &Nodes{}
	heap.Init(h)
	visitedNodes := map[Point]bool{}
	heap.Push(h, Node{p: start, steps: 0})
	i := 0
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		if _, ok := visitedNodes[n.p]; ok {
			continue
		} else {
			visitedNodes[n.p] = true
		}
		if n.p.x == end.x && n.p.y == end.y {
			return n.steps
		}
		points, _ := findAscNextSteps(n.p, m)
		for _, e := range points {
			heap.Push(h, Node{p: e, steps: n.steps + 1})
		}
		i++
	}
	return 0
}

func solvePart2(m [][]rune, start Point) int {
	h := &Nodes{}
	heap.Init(h)
	visitedNodes := map[Point]bool{}
	heap.Push(h, Node{p: start, steps: 0})
	i := 0
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		if _, ok := visitedNodes[n.p]; ok {
			continue
		} else {
			visitedNodes[n.p] = true
		}
		if m[n.p.y][n.p.x] == 'a' {
			return n.steps
		}

		points, _ := findDescNextSteps(n.p, m)
		for _, e := range points {
			heap.Push(h, Node{p: e, steps: n.steps + 1})
		}
		i++
	}
	return 0
}

func findAscNextSteps(p Point, m [][]rune) ([]Point, []int) {
	out := []Point{}
	h := []int{}
	for _, d := range DIRECTIONS {
		nextX := p.x + d[0]
		nextY := p.y + d[1]
		if len(m[0]) > nextX && nextX >= 0 && len(m) > nextY && nextY >= 0 {
			loc := m[nextY][nextX]
			height := loc - m[p.y][p.x]
			if height <= 1 {
				out = append(out, Point{x: nextX, y: nextY})
				h = append(h, int(height))
			}
		}
	}
	return out, h
}

func findDescNextSteps(p Point, m [][]rune) ([]Point, []int) {
	out := []Point{}
	h := []int{}
	for _, d := range DIRECTIONS {
		nextX := p.x + d[0]
		nextY := p.y + d[1]
		if len(m[0]) > nextX && nextX >= 0 && len(m) > nextY && nextY >= 0 {
			loc := m[nextY][nextX]
			height := m[p.y][p.x] - loc
			if height <= 1 {
				out = append(out, Point{x: nextX, y: nextY})
				h = append(h, int(height))
			}
		}
	}
	return out, h
}

func find(r rune, m [][]rune) Point {
	for y, rows := range m {
		for x, col := range rows {
			if col == r {
				return Point{
					x: x,
					y: y,
				}
			}
		}
	}
	return Point{}
}

func parse(input string) [][]rune {
	out := [][]rune{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		out = append(out, []rune(row))
	}
	return out
}

func printMatrix(matrix [][]rune, visited []Point) {
	for y, m := range matrix {
		r := ""
		for x, c := range m {
			l := string(c)
			for _, v := range visited {
				if v.x == x && v.y == y {
					l = Red + l + Reset
				}
			}
			r += l
		}
		fmt.Println(r)
	}
}
