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

type Point struct {
	x int
	y int
	w int
}

type JunctionNode struct {
	point Point
	next  map[Point]int
}

type Node struct {
	point   Point
	steps   int
	visited []Point
}

type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].steps > h[j].steps }
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
func FindNodes(sx, sy, ex, ey int, m [][]rune) []*JunctionNode {
	out := []*JunctionNode{{
		point: Point{x: sx, y: sy},
		next:  map[Point]int{}}, {
		point: Point{x: ex, y: ey},
		next:  map[Point]int{}}}
	for y, row := range m {
		for x := range row {
			n := FindNextStepsPart2(x, y, 0, make(map[string]int), m)
			if len(n) > 2 {
				out = append(out, &JunctionNode{
					point: Point{x: x, y: y}, next: map[Point]int{}})
			}
		}
	}
	return out

}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input, 139, 140)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input, 139, 140)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string, ex, ey int) int {
	s := []int{1, 0}
	m := Parse(input)
	steps := solveTask1(s[0], s[1], ex, ey, m)

	return steps
}

func task2(input string, ex, ey int) int {
	s := []int{1, 0}
	m := Parse(input)
	n := FindNodes(s[0], s[1], ex, ey, m)
	for _, e := range n {
		e.ProcessNodes(n, m)
	}
	return solveTask2(ex, ey, n, m)
}

func (n *JunctionNode) ProcessNodes(nodes []*JunctionNode, m [][]rune) {
	queue := make([]Point, 0)
	visited := make(map[string]int)
	key := fmt.Sprintf("x:%d,y%d", n.point.x, n.point.y)
	visited[key] = 0

	nextPoints := FindNextStepsPart2(n.point.x, n.point.y, 0, visited, m)
	queue = append(queue, nextPoints...)

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		key = fmt.Sprintf("x:%d,y%d", next.x, next.y)
		visited[key] = 0

		if isNode(next.x, next.y, nodes) {
			n.next[Point{x: next.x, y: next.y}] = next.w
		} else {
			queue = append(queue, FindNextStepsPart2(next.x, next.y, next.w, visited, m)...)
		}
	}
}

func isNode(x, y int, nodes []*JunctionNode) bool {
	for _, n := range nodes {
		if n.point.x == x && n.point.y == y {
			return true
		}
	}
	return false
}

func solveTask1(sx, sy, ex, ey int, m [][]rune) int {
	h := &NodeHeap{}
	maxSteps := 0
	heap.Init(h)
	heap.Push(h, Node{point: Point{sx, sy, 1}, steps: 0, visited: []Point{}})
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		v := []Point{{x: n.point.x, y: n.point.y}}
		v = append(v, n.visited...)
		if n.point.x == ex && n.point.y == ey {
			if n.steps > maxSteps {
				maxSteps = n.steps
			}
		}
		if Contains(n.point, n.visited) {
			continue
		}
		nextNodes := FindNextStepsPart1(n.point.x, n.point.y, m)
		for _, e := range nextNodes {
			heap.Push(h, Node{point: e, steps: n.steps + 1, visited: v})
		}
	}
	return maxSteps
}

func solveTask2(ex, ey int, nodes []*JunctionNode, m [][]rune) int {
	queue := []Node{}
	maxSteps := 0
	queue = append(queue, Node{point: nodes[0].point, steps: 0, visited: []Point{}})
	for len(queue) > 0 {
		next := queue[0]
		v := []Point{{x: next.point.x, y: next.point.y, w: next.steps}}
		v = append(v, next.visited...)
		queue = queue[1:]
		if Contains(next.point, next.visited) {
			continue
		}
		if next.point.x == ex && next.point.y == ey {
			if next.steps > maxSteps {
				maxSteps = next.steps
			}
		} else {
			n := getNode(next.point, nodes)
			for k, s := range n.next {
				queue = append(queue, Node{point: Point{x: k.x, y: k.y}, steps: next.steps + s, visited: v})
			}
		}
	}
	return maxSteps
}

func getNode(point Point, nodes []*JunctionNode) *JunctionNode {
	for _, v := range nodes {
		if v.point.x == point.x && v.point.y == point.y {
			return v
		}
	}
	return nil
}

func Parse(input string) [][]rune {
	out := [][]rune{}
	row := strings.Split(input, "\n")
	for _, r := range row {
		out = append(out, []rune(r))
	}
	return out
}

func Contains(curr Point, visited []Point) bool {
	for _, v := range visited {
		if v.x == curr.x && v.y == curr.y {
			return true
		}
	}
	return false
}

func FindNextStepsPart2(x, y, w int, visited map[string]int, m [][]rune) []Point {
	nextSteps := []Point{}
	possibleSteps := [][]int{SOUTH, NORTH, WEST, EAST}
	for _, s := range possibleSteps {
		nx := x + s[0]
		ny := y + s[1]
		key := fmt.Sprintf("x:%d,y%d", nx, ny)
		_, ok := visited[key]
		if ok {
			continue
		}
		if nx < 0 || nx >= len(m[0]) || ny < 0 || ny >= len(m) {
			continue
		}
		t := m[ny][nx]
		if t != '#' {
			nextSteps = append(nextSteps, Point{nx, ny, w + 1})
		}
	}
	return nextSteps
}

func FindNextStepsPart1(x, y int, m [][]rune) []Point {
	nextSteps := []Point{}
	slopes := []rune{'v', '^', '<', '>'}
	possibleSteps := [][]int{SOUTH, NORTH, WEST, EAST}
	ct := m[y][x]
	for i, v := range slopes {
		if v == ct {
			nx := x + possibleSteps[i][0]
			ny := y + possibleSteps[i][1]
			if nx >= 0 && nx < len(m[0]) && ny >= 0 || ny < len(m) {
				return []Point{{nx, ny, 1}}
			}
		}
	}

	for _, s := range possibleSteps {
		nx := x + s[0]
		ny := y + s[1]
		if nx < 0 || nx >= len(m[0]) || ny < 0 || ny >= len(m) {
			continue
		}
		t := m[ny][nx]
		if t != '#' {
			nextSteps = append(nextSteps, Point{nx, ny, 1})
		}
	}
	return nextSteps
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
