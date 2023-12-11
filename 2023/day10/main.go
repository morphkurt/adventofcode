package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Location struct {
	x    int
	y    int
	Type string
}

type Point struct {
	x int
	y int
}

var NORTH = [2]int{0, -1}
var SOUTH = [2]int{0, 1}
var WEST = [2]int{-1, 0}
var EAST = [2]int{1, 0}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	r, m := ReadMap(input)
	x, y := FindOrigin(r)
	d := solve(r, m, Point{x: -1, y: -1}, Point{x: x, y: y})
	return ((d - 1) / 2)
}

func task2(input string) int {
	r, m := ReadMap(input)
	x, y := FindOrigin(r)
	m[y][x] = "S"
	solve(r, m, Point{x: -1, y: -1}, Point{x: x, y: y})
	e := Expand(m)
	Drain(e)
	s := Shrink(e)
	return Count(s)
}

func solve(m [][]string, m1 [][]string, p Point, c Point) (distance int) {
	if m[c.y][c.x] == "S" && p.x >= 0 {
		return 1
	} else {
		loc := GetNextLocation(m, c.x, c.y)
		for _, l := range loc {
			if l.x != p.x || l.y != p.y {
				m1[l.y][l.x] = l.Type
				d := solve(m, m1, c, Point{x: l.x, y: l.y})
				if d > 0 {
					return d + 1
				}
				m1[l.y][l.x] = "0"
			}
		}
	}
	return 0
}

func ReadMap(input string) (Map [][]string, marker [][]string) {
	result := [][]string{}
	marker = [][]string{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		r := []string{}
		result = append(result, s)
		for i := 0; i < len(s); i++ {
			r = append(r, "0")
		}
		marker = append(marker, r)
	}
	return result, marker
}

func FindOrigin(m [][]string) (x, y int) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == "S" {
				return x, y
			}
		}
	}
	return -1, -1
}

func GetNextLocation(m [][]string, x, y int) (loc []Location) {
	loc = []Location{}
	s := m[y][x]
	n := [][]int{}
	switch s {
	case "F":
		n = [][]int{
			SOUTH[:], //south
			EAST[:],  //east
		}
	case "|":
		n = [][]int{
			NORTH[:], //north
			SOUTH[:], //south
		}
	case "L":
		n = [][]int{
			NORTH[:], //north
			EAST[:],  //east
		}
	case "J":
		n = [][]int{
			NORTH[:], //north
			WEST[:],  //west
		}
	case "7":
		n = [][]int{
			WEST[:],  //west
			SOUTH[:], //south
		}
	case "S":
		n = [][]int{
			NORTH[:], //north
			WEST[:],  //west
			SOUTH[:], //south
			EAST[:],  //east
		}
	case "-":
		n = [][]int{
			WEST[:], //west
			EAST[:], //east
		}
	}
	for _, e := range n {
		if x+e[0] >= 0 && x+e[0] < len(m[0]) && y+e[1] >= 0 && y+e[1] < len(m) {
			v := m[y+e[1]][x+e[0]]
			if v != "." {
				loc = append(loc, Location{
					x:    x + e[0],
					y:    y + e[1],
					Type: v,
				})
			}

		}
	}
	return loc
}

func PrintMap(mark [][]string) {
	for _, r := range mark {
		fmt.Println(strings.Join(r, ""))
	}
}

func Shrink(mark [][]string) (out [][]string) {
	out = [][]string{}
	for y := 0; y < len(mark); y = y + 2 {
		row := []string{}
		for x := 0; x < len(mark[0]); x = x + 2 {
			row = append(row, mark[y][x])
		}
		out = append(out, row)
	}
	return out
}

func Count(mark [][]string) (out int) {
	count := 0
	for y := 0; y < len(mark); y++ {
		for x := 0; x < len(mark[0]); x++ {
			if mark[y][x] == "0" {
				count++
			}
		}
	}
	return count
}

func Expand(mark [][]string) (out [][]string) {
	n := [][]int{
		SOUTH[:],
		EAST[:],
	}
	out = [][]string{}
	for y, row := range mark {
		r := []string{}
		r1 := []string{}
		for x, v := range row {
			switch v {
			case "0":
				r = append(r, "0", "0")
				r1 = append(r1, "0", "0")
			case "S":
				for i, e := range n {
					if x+e[0] >= 0 && x+e[0] < len(mark[0]) && y+e[1] >= 0 && y+e[1] < len(mark) {
						p := mark[y+e[1]][x+e[0]]
						if i == 0 { // south
							if slices.Contains([]string{"|", "J", "L"}, p) {
								r1 = append(r1, "|", "0")
							}
							if slices.Contains([]string{"-", "F", "7", "O"}, p) {
								r1 = append(r1, "0", "0")
							}
						} else { // east
							if slices.Contains([]string{"|", "F", "L", "S"}, p) {
								r = append(r, "S", "0")
							}
							if slices.Contains([]string{"-", "7", "J", "O"}, p) {
								r = append(r, "S", "-")
							}
						}
					} else {
						if i == 0 { // south
							r1 = append(r1, "0", "0")

						} else {
							r = append(r, "S", "0")

						}
					}
				}
			case "|":
				r = append(r, "|", "0")
				r1 = append(r1, "|", "0")
			case "-":
				r = append(r, "-", "-")
				r1 = append(r1, "0", "0")
			case "7":
				r = append(r, "7", "0")
				r1 = append(r1, "|", "0")
			case "J":
				r = append(r, "J", "0")
				r1 = append(r1, "0", "0")
			case "F":
				r = append(r, "F", "-")
				r1 = append(r1, "|", "0")
			case "L":
				r = append(r, "L", "-")
				r1 = append(r1, "0", "0")
			}
		}
		out = append(out, r, r1)
	}
	return out
}

func Drain(mark [][]string) {
	for y := 0; y < len(mark); y++ {
		MarkDrained(mark, Point{x: 0, y: y})
		MarkDrained(mark, Point{x: len(mark[0]) - 1, y: y})
	}
	for x := 0; x < len(mark[0]); x++ {
		MarkDrained(mark, Point{x: x, y: 0})
		MarkDrained(mark, Point{x: x, y: len(mark) - 1})
	}
}

func MarkDrained(mark [][]string, p Point) {
	if mark[p.y][p.x] == "0" {
		mark[p.y][p.x] = "#"
		n := GetEmptyNeighours(mark, p)
		if len(n) == 0 {
			return
		} else {
			for _, e := range n {
				MarkDrained(mark, e)
			}
		}
	}

}

func GetEmptyNeighours(mark [][]string, p Point) (points []Point) {
	points = []Point{}
	n := [][]int{
		NORTH[:], //north
		WEST[:],  //west
		SOUTH[:], //south
		EAST[:],  //east
	}
	for _, e := range n {
		if p.x+e[0] >= 0 && p.x+e[0] < len(mark[0]) && p.y+e[1] >= 0 && p.y+e[1] < len(mark) {
			if mark[p.y+e[1]][p.x+e[0]] == "0" {
				points = append(points, Point{x: p.x + e[0],
					y: p.y + e[1]})
			}
		}
	}
	return points
}
