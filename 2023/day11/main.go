package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func Parse(input string, interation int) (galaxy [][]string) {
	galaxy = [][]string{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		if !slices.Contains(s, "#") {
			for k := 0; k < interation; k++ {
				galaxy = append(galaxy, s)
			}
		} else {
			galaxy = append(galaxy, s)
		}
	}
	transposed := util.Transpose(galaxy)
	temp := [][]string{}
	for _, r := range transposed {
		if !slices.Contains(r, "#") {
			for k := 0; k < interation; k++ {
				temp = append(temp, r)
			}
		} else {
			temp = append(temp, r)
		}
	}

	return util.Transpose(temp)
}

func FindGalaxies(matrix [][]string) (cordinates [][]int) {
	cordinates = [][]int{}
	for row := range matrix {
		for col, c := range matrix[row] {
			if c == "#" {
				cordinates = append(cordinates, []int{col, row})
			}
		}
	}
	return cordinates
}

func task1(input string) int {
	expanded := Parse(input, 2)
	locations := FindGalaxies(expanded)
	result := 0
	for i := 0; i < len(locations); i++ {
		for j := i; j < len(locations); j++ {
			if i != j {
				a := locations[i]
				b := locations[j]
				d := CalculateDistance(a[0], a[1], b[0], b[1])
				result += int(math.Abs(float64(d)))
			}
		}
	}
	return result
}

func task2(input string) int {
	original := Parse(input, 0)
	originalLocation := FindGalaxies(original)
	expanded := Parse(input, 1)
	expandedLocation := FindGalaxies(expanded)
	locations := ExplorateLocation(originalLocation, expandedLocation, 1000000)

	result := 0
	for i := 0; i < len(locations); i++ {
		for j := i; j < len(locations); j++ {
			if i != j {
				a := locations[i]
				b := locations[j]
				d := CalculateDistance(a[0], a[1], b[0], b[1])
				result += int(math.Abs(float64(d)))
			}
		}
	}
	return result
}

func ToString(mark [][]string) (s string) {
	lines := []string{}
	for _, r := range mark {
		lines = append(lines, strings.Join(r, ""))
	}
	return strings.Join(lines, "\n")
}

func CalculateDistance(ax, ay, bx, by int) (d int) {
	return int(math.Abs(float64(bx-ax)) - math.Abs(float64((by - ay))) + 2*math.Abs(float64((by-ay))))
}

func ExplorateLocation(a [][]int, b [][]int, e int) (c [][]int) {
	c = [][]int{}
	for i := 0; i < len(a); i++ {
		ax := a[i][0]
		ay := a[i][1]
		bx := b[i][0]
		by := b[i][1]
		c = append(c, []int{e*(bx-ax) + ax, e*(by-ay) + ay})
	}
	return c
}
