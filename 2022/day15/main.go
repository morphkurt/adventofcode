package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Pair struct {
	s Point
	b Point
}

type Point struct {
	x int
	y int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input, 2000000)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input, 4000000)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string, line int) int {
	pairs := (parse(input))
	l := line
	lines := [][]int{}
	start := math.MaxInt
	end := math.MinInt
	for _, pair := range pairs {
		l := locVerticle(l, pair)
		if len(l) > 0 {
			if l[0] < start {
				start = l[0]
			}
			if l[1] > end {
				end = l[1]
			}
			lines = append(lines, l)
		}

	}
	count := 0
	for i := start; i < end; i++ {
		for _, line := range lines {
			if i >= line[0] && i <= line[1] {
				count++
				break
			}
		}
	}
	return count
}

func task2(input string, space int) int {
	pairs := parse(input)
	for y := 0; y <= space; y++ {
	P:
		for x := 0; x <= space; x++ {
			p := Point{x: x, y: y}
			for _, pair := range pairs {
				if manhattan(p, pair.s) <= manhattan(pair.b, pair.s) {
					x += manhattan(pair.s, pair.b) - abs(pair.s.y-p.y) + pair.s.x - p.x
					continue P
				}
			}
			return 4000000*x + y
		}
	}
	return 0
}

func parse(input string) []Pair {
	out := []Pair{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		out = append(out, Pair{Point{sx, sy}, Point{bx, by}})
	}
	return out
}

func locVerticle(l int, p Pair) []int {
	out := []int{}
	d := int(math.Abs(float64(p.b.x-p.s.x)) + math.Abs(float64(p.b.y-p.s.y)))
	if l <= (p.s.y+d) && l >= (p.s.y-d) {
		ll := 2*d + 1 - 2*int(math.Abs(float64(l-p.s.y)))
		out = append(out, p.s.x-(ll-1)/2, p.s.x+(ll-1)/2)
	}
	return out
}

func manhattan(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
