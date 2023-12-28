package main

import (
	"fmt"
	"sort"
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

func task1(input string) int {
	patch := parse(input)
	visibleCount := 0
	for y := 0; y < len(patch); y++ {
		for x := 0; x < len(patch[0]); x++ {
			if y == 0 || x == 0 || y == len(patch)-1 || x == len(patch[0])-1 {
				visibleCount++
			} else {
				visible := []bool{true, true, true, true}
				for n := 0; n < y; n++ {
					if patch[n][x] >= patch[y][x] {
						visible[0] = false
						break
					}
				}
				for n := y + 1; n < len(patch); n++ {
					if patch[n][x] >= patch[y][x] {
						visible[1] = false
						break
					}
				}
				for n := 0; n < x; n++ {
					if patch[y][n] >= patch[y][x] {
						visible[2] = false
						break
					}
				}
				for n := x + 1; n < len(patch[0]); n++ {
					if patch[y][n] >= patch[y][x] {
						visible[3] = false
						break
					}
				}
				for _, v := range visible {
					if v == true {
						visibleCount++
						break
					}
				}
			}
		}
	}
	return visibleCount
}

func task2(input string) int {
	patch := parse(input)
	scenicScore := []int{}
	for y := 0; y < len(patch); y++ {
		for x := 0; x < len(patch[0]); x++ {
			dirCount := []int{0, 0, 0, 0}
			for n := y - 1; n >= 0; n-- {
				dirCount[0]++
				if patch[n][x] >= patch[y][x] {
					break
				}
			}
			for n := y + 1; n < len(patch); n++ {
				dirCount[1]++
				if patch[n][x] >= patch[y][x] {
					break
				}
			}
			for n := x - 1; n >= 0; n-- {
				dirCount[2]++
				if patch[y][n] >= patch[y][x] {
					break
				}
			}
			for n := x + 1; n < len(patch[0]); n++ {
				dirCount[3]++
				if patch[y][n] >= patch[y][x] {
					break
				}
			}
			scenicScore = append(scenicScore, dirCount[0]*dirCount[1]*dirCount[2]*dirCount[3])
		}
	}
	sort.Ints(scenicScore)
	return scenicScore[len(scenicScore)-1]
}

func parse(input string) [][]int {
	out := [][]int{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		c := strings.Split(row, "")
		out = append(out, util.ToInt(c))
	}
	return out
}
