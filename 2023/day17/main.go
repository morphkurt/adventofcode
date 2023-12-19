package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var minWeight = math.MaxInt64

var START = []int{0, 0, 0}
var NORTH = []int{1, 0, -1}
var SOUTH = []int{2, 0, 1}
var EAST = []int{3, 1, 0}
var WEST = []int{4, -1, 0}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type Point struct {
	x int
	y int
}

type Node struct {
	x         int
	y         int
	weight    int
	steps     int
	direction []int // how did we get here
	visited   []Point
}

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *NodeHeap) Pop() interface{} {

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
	m := Parse(input)
	return solveTask1(m)
}

func task2(input string) int {
	m := Parse(input)
	return solveTask2(m)
}

func Parse(input string) (matrix [][]int) {
	matrix = [][]int{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		r := []int{}
		for _, t := range row {
			r = append(r, int(t-'0'))
		}
		matrix = append(matrix, r)
	}
	return matrix
}

func Contains(c []Point, e Point) bool {
	for _, p := range c {
		if p.x == e.x && p.y == e.y {
			return true
		}
	}
	return false
}

func solveTask2(m [][]int) int {
	h := &NodeHeap{}
	//w := 0
	heap.Init(h)
	n := getNodesTask2(0, 0, 0, 0, START, m, []Point{})
	var cache = make(map[string]int)
	// pushing the first two nodes
	for _, e := range n {
		heap.Push(h, e)
	}
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		v := []Point{{x: n.x, y: n.y}}
		v = append(v, n.visited...)

		key := fmt.Sprintf("%d:%d:%d,%d", n.x, n.y, n.direction[0], n.steps)
		if val, ok := cache[key]; ok && val <= n.weight {
			continue
		}
		cache[key] = n.weight
		if n.x == len(m[0])-1 && n.y == len(m)-1 {
			return n.weight
		}
		if Contains(n.visited, Point{x: n.x, y: n.y}) {
			continue
		}
		nextNodes := getNodesTask2(n.x, n.y, n.steps, n.weight, n.direction, m, v)
		for _, e := range nextNodes {
			heap.Push(h, e)
		}
	}
	return -1

}

func solveTask1(m [][]int) int {
	h := &NodeHeap{}
	//w := 0
	heap.Init(h)
	n := getNodesTask1(0, 0, 0, 0, START, m, []Point{})
	var cache = make(map[string]int)
	// pushing the first two nodes
	for _, e := range n {
		heap.Push(h, e)
	}
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		v := []Point{{x: n.x, y: n.y}}
		v = append(v, n.visited...)
		key := fmt.Sprintf("%d:%d:%d,%d", n.x, n.y, n.direction[0], n.steps)
		if val, ok := cache[key]; ok && val <= n.weight {
			continue
		}
		cache[key] = n.weight
		if n.x == len(m[0])-1 && n.y == len(m)-1 {
			return n.weight
		}
		if Contains(n.visited, Point{x: n.x, y: n.y}) {
			continue
		}
		nextNodes := getNodesTask1(n.x, n.y, n.steps, n.weight, n.direction, m, v)
		for _, e := range nextNodes {
			heap.Push(h, e)
		}
	}
	return -1
}

func getNodesTask1(x, y, d, prevW int, direction []int, matrix [][]int, visited []Point) []Node {
	nodes := []Node{}
	if d > 0 {
		nextX := x + direction[1]
		nextY := y + direction[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: direction, steps: d - 1, visited: visited})
		}
	}
	switch direction[0] {
	case START[0]:
		// next node coordinates
		nextX := x + SOUTH[1]
		nextY := y + SOUTH[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: SOUTH, steps: 2, visited: visited})
		}
		nextX = x + EAST[1]
		nextY = y + EAST[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: EAST, steps: 2, visited: visited})
		}
	case NORTH[0], SOUTH[0]:
		// no restriction of moving
		nextX := x + EAST[1]
		nextY := y + EAST[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: EAST, steps: 2, visited: visited})
		}
		nextX = x + WEST[1]
		nextY = y + WEST[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: WEST, steps: 2, visited: visited})
		}
	case EAST[0], WEST[0]:
		// no restriction of moving
		nextX := x + NORTH[1]
		nextY := y + NORTH[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: NORTH, steps: 2, visited: visited})
		}
		nextX = x + SOUTH[1]
		nextY = y + SOUTH[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: SOUTH, steps: 2, visited: visited})
		}
	}
	return nodes
}

func getNodesTask2(x, y, d, prevW int, direction []int, matrix [][]int, visited []Point) []Node {
	minDistance := 7 // 9,8,7
	destRadius := 3
	nodes := []Node{}
	if d > 0 {
		nextX := x + direction[1]
		nextY := y + direction[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: direction, steps: d - 1, visited: visited})
		}
	}
	switch direction[0] {
	case START[0]:
		// next node coordinates
		nextX := x + SOUTH[1]
		nextY := y + SOUTH[2]
		nextIsDestination := (len(matrix)-1-nextY >= destRadius)
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance && nextIsDestination {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: SOUTH, steps: 9, visited: visited})
		}
		nextX = x + EAST[1]
		nextY = y + EAST[2]
		nextIsDestination = (len(matrix[0])-1-nextX >= destRadius)
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance && nextIsDestination {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: EAST, steps: 9, visited: visited})
		}
	case NORTH[0], SOUTH[0]:
		// no restriction of moving
		nextX := x + EAST[1]
		nextY := y + EAST[2]
		nextIsDestination := (len(matrix[0])-1-nextX >= destRadius)
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance && nextIsDestination {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: EAST, steps: 9, visited: visited})
		}
		nextX = x + WEST[1]
		nextY = y + WEST[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: WEST, steps: 9, visited: visited})
		}
	case EAST[0], WEST[0]:
		// no restriction of moving
		nextX := x + NORTH[1]
		nextY := y + NORTH[2]
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: NORTH, steps: 9, visited: visited})
		}
		nextX = x + SOUTH[1]
		nextY = y + SOUTH[2]
		nextIsDestination := (len(matrix)-1-nextY >= destRadius)
		if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) && d < minDistance && nextIsDestination {
			w := matrix[nextY][nextX]
			nodes = append(nodes, Node{x: nextX, y: nextY, weight: w + prevW, direction: SOUTH, steps: 9, visited: visited})
		}
	}
	return nodes
}

func printMatrix(matrix [][]int, visited []Point) {
	for y, m := range matrix {
		r := ""
		for x, c := range m {
			l := strconv.Itoa(c)
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
