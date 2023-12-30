package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var UP = []int{0, -1}
var DOWN = []int{0, 1}
var LEFT = []int{-1, 0}
var RIGHT = []int{1, 0}

var NE = []int{1, -1}
var NW = []int{-1, -1}
var SE = []int{1, 1}
var SW = []int{-1, 1}

var DIRECTION = map[string][]int{
	"U": UP,
	"D": DOWN,
	"L": LEFT,
	"R": RIGHT,
}

var DIAGONAL = map[string][]int{
	"NE": NE,
	"NW": NW,
	"SE": SE,
	"SW": SW,
}

type Point struct {
	x int
	y int
}

type Rope struct {
	knots [][]Point
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	rope, _ := parse(input)
	return countLocations(rope.knots[0])
}

func task2(input string) int {
	_, rope := parse(input)
	return countLocations(rope.knots[0])
}

func countLocations(points []Point) int {
	loc := map[Point]int{}
	for _, v := range points {
		if _, ok := loc[v]; ok {
			loc[v] = loc[v] + 1
		} else {
			loc[v] = 1
		}
	}
	return len(loc)
}

func parse(input string) (Rope, Rope) {
	lines := strings.Split(input, "\n")
	knots := [][]Point{}
	for i := 0; i < 2; i++ {
		knots = append(knots, []Point{Point{0, 0}})
	}
	twoKnotRope := Rope{knots: knots}
	knots = [][]Point{}
	for i := 0; i < 10; i++ {
		knots = append(knots, []Point{Point{0, 0}})
	}
	tenKnotRope := Rope{knots: knots}

	for _, line := range lines {
		var direction string
		var distance int
		fmt.Sscanf(line, "%s %d", &direction, &distance)
		twoKnotRope.move(DIRECTION[direction], distance)
		tenKnotRope.move(DIRECTION[direction], distance)

	}
	return twoKnotRope, tenKnotRope
}

func (rope *Rope) move(direction []int, distance int) {
	knots := rope.knots
	for steps := 1; steps <= distance; steps++ {
		currHeadX := (knots[len(knots)-1][len(knots[len(knots)-1])-1].x)
		currHeadY := (knots[len(knots)-1][len(knots[len(knots)-1])-1].y)
		rope.knots[len(knots)-1] = append(rope.knots[len(knots)-1], Point{(currHeadX) + (direction[0]), (currHeadY) + (direction[1])})
		for k := len(rope.knots) - 1; k > 0; k-- {
			curKnot := rope.knots[k]
			prevKnot := rope.knots[k-1]
			nextStep, shouldMove := nextStep(prevKnot[len(prevKnot)-1], curKnot[len(curKnot)-1])
			if shouldMove {
				rope.knots[k-1] = append(rope.knots[k-1], nextStep)
			}
		}
	}
}

func nextStep(knot1 Point, knot2 Point) (Point, bool) {
	absDistance := math.Max(math.Abs(float64(knot1.x-knot2.x)),
		math.Abs(float64(knot1.y-knot2.y)))
	if absDistance <= 1 {
		return knot1, false
	}
	for _, v := range DIRECTION {
		//check whether two steps U/D/L/R
		if knot2.x == knot1.x+2*v[0] && knot2.y == knot1.y+2*v[1] {
			return Point{knot1.x + v[0], knot1.y + v[1]}, true
		}
	}
	if knot2.x > knot1.x && knot2.y > knot1.y {
		return Point{knot1.x + SE[0], knot1.y + SE[1]}, true
	}
	if knot2.x > knot1.x && knot2.y < knot1.y {
		return Point{knot1.x + NE[0], knot1.y + NE[1]}, true
	}
	if knot2.x < knot1.x && knot2.y < knot1.y {
		return Point{knot1.x + NW[0], knot1.y + NW[1]}, true
	}
	if knot2.x < knot1.x && knot2.y > knot1.y {
		return Point{knot1.x + SW[0], knot1.y + SW[1]}, true
	}

	return knot1, false

}

func printRope(rope Rope) {
	tailLocations := rope.knots[0]
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for _, v := range tailLocations {
		if v.x < minX {
			minX = v.x
		}
		if v.y < minY {
			minY = v.y
		}
		if v.x > maxX {
			maxX = v.x
		}
		if v.y > maxY {
			maxY = v.y
		}
	}
	loc := [][]rune{}
	for y := minY; y <= maxY; y++ {
		row := []rune{}
		for x := minX; x <= maxX; x++ {
			row = append(row, '.')
		}
		loc = append(loc, row)
	}
	for _, v := range tailLocations {
		loc[v.y-minY][v.x-minX] = '#'
	}
	for _, row := range loc {
		fmt.Println(string(row))
	}
}
