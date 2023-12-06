package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
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
	result := 1
	times, distances := parse(input)
	for i, t := range times {
		r := 0
		d := distances[i]
		for j := 0; j <= t; j++ {
			if (t-j)*j > d {
				r++
			}
		}
		result = result * r
	}
	return result

}

func task2(input string) int {
	times, distances := parse(input)
	concatTimes := ""
	concatDistances := ""
	for i, t := range times {
		concatTimes += strconv.Itoa(t)
		concatDistances += strconv.Itoa(distances[i])
	}
	t, _ := strconv.Atoi(concatTimes)
	d, _ := strconv.Atoi(concatDistances)

	r := 0
	for j := 0; j <= t; j++ {
		if (t-j)*j > d {
			r++
		}
	}
	return r

}

func parse(input string) ([]int, []int) {
	time, distance := []int{}, []int{}
	var re = regexp.MustCompile(`(?m)(\d)+`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		ds := re.FindAllString(line, -1)
		for _, d := range ds {
			dv, _ := strconv.Atoi(d)
			if strings.Contains(line, "Time") {
				time = append(time, dv)
			}
			if strings.Contains(line, "Distance") {
				distance = append(distance, dv)
			}
		}
	}
	return time, distance
}
