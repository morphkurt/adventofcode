package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var D = []int{0, 1}
var L = []int{-1, 0}
var R = []int{1, 0}

type Point struct {
	x int
	y int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	p := parse(input)
	m, _, _, minX, _ := toMapTask1(p)
	c := 0
	i := 1
	for {
		dropSand(m, Point{500 - minX, 0})
		cc := count(m)
		if cc > c {
			c = cc
		} else {
			return i - 1
		}
		i++
	}
}

func task2(input string) int {
	p := parse(input)
	m, _, _, _, _ := toMapTask2(p)
	c := 0
	i := 0
	for {
		dropSand(m, Point{len(m[0])/2 + 1, 0})
		cc := count(m)
		if cc > c {
			c = cc
		} else {
			return i
		}
		i++
	}
}

func parse(input string) [][]Point {
	out := [][]Point{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lc := []Point{}
		points := strings.Split(line, " -> ")
		for _, p := range points {
			c := util.ToInt(strings.Split(p, ","))
			lc = append(lc, Point{c[0], c[1]})
		}
		out = append(out, lc)
	}
	return out
}

func nextMove(m [][]rune, s Point) (Point, bool) {
	moves := []Point{
		{x: s.x + D[0], y: s.y + D[1]},               //down
		{x: s.x + D[0] + L[0], y: s.y + D[1] + L[1]}, //down left
		{x: s.x + D[0] + R[0], y: s.y + D[1] + R[1]}, //down right
	}
	for _, move := range moves {
		if within(m, move) {
			loc := m[move.y][move.x]
			if loc == '.' {
				return move, false
			}
		} else {
			return s, true
		}
	}
	return s, false
}

func within(m [][]rune, s Point) bool {
	if s.y >= len(m) || s.y < 0 {
		return false
	}
	if s.x >= len(m[0]) || s.x < 0 {
		return false
	}
	return true
}

func dropSand(m [][]rune, s Point) {
	settled := false
	hitAbysis := false
	next, curr := s, s
	for {
		next, hitAbysis = nextMove(m, curr)
		if curr.x == next.x && curr.y == next.y {
			settled = true
		}
		if hitAbysis {
			break
		}
		if settled {
			m[next.y][next.x] = 'o'
			break
		}
		curr = next
	}
}

func toMapTask1(p [][]Point) (out [][]rune, maxX, maxY, minX, minY int) {
	out = [][]rune{}
	maxX, maxY = math.MinInt, math.MinInt
	minX, minY = math.MaxInt, 0
	for _, lines := range p {
		for _, point := range lines {
			if point.x > maxX {
				maxX = point.x
			}
			if point.y > maxY {
				maxY = point.y
			}
			if point.x < minX {
				minX = point.x
			}
		}
	}
	for y := 0; y <= maxY-minY; y++ {
		row := []rune{}
		for x := 0; x <= maxX-minX; x++ {
			row = append(row, '.')
		}
		out = append(out, row)
	}
	for _, lines := range p {
		for i := 0; i < len(lines)-1; i++ {
			cur := lines[i]
			next := lines[i+1]
			fromX := math.Min(float64(cur.x), float64(next.x))
			toX := math.Max(float64(cur.x), float64(next.x))
			fromY := math.Min(float64(cur.y), float64(next.y))
			toY := math.Max(float64(cur.y), float64(next.y))
			for y := fromY; y <= toY; y++ {
				for x := fromX; x <= toX; x++ {
					out[int(y)-minY][int(x)-minX] = '#'
				}
			}
		}
	}
	out[0][500-minX] = '+'
	return out, maxX, maxY, minX, minY
}

func toMapTask2(p [][]Point) (out [][]rune, maxX, maxY, minX, minY int) {
	out = [][]rune{}
	maxX, maxY = math.MinInt, math.MinInt
	minX, minY = math.MaxInt, 0
	for _, lines := range p {
		for _, point := range lines {
			if point.x > maxX {
				maxX = point.x
			}
			if point.y > maxY {
				maxY = point.y
			}
			if point.x < minX {
				minX = point.x
			}
		}
	}
	for y := 0; y <= (maxY-minY)+2; y++ {
		row := []rune{}
		for x := 0; x <= (maxY-minY+3)*2; x++ {
			row = append(row, '.')
		}
		out = append(out, row)
	}
	for _, lines := range p {
		for i := 0; i < len(lines)-1; i++ {
			cur := lines[i]
			next := lines[i+1]
			offset := (len(out[0])/2 + 1) - (500 - minX)
			fromX := math.Min(float64(cur.x), float64(next.x))
			toX := math.Max(float64(cur.x), float64(next.x))
			fromY := math.Min(float64(cur.y), float64(next.y))
			toY := math.Max(float64(cur.y), float64(next.y))
			for y := fromY; y <= toY; y++ {
				for x := fromX; x <= toX; x++ {
					out[int(y)-minY][int(x)-minX+offset] = '#'
				}
			}
		}
	}
	for x := 0; x < len(out[0]); x++ {
		out[len(out)-1][x] = '#'
	}
	out[0][len(out[0])/2+1] = '+'
	return out, maxX, maxY, minX, minY
}

func printMatrix(m [][]rune) {
	for _, line := range m {
		fmt.Println(string(line))
	}
}

func count(m [][]rune) int {
	count := 0
	for _, row := range m {
		for _, c := range row {
			if c == 'o' {
				count++
			}
		}
	}
	return count
}
