package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var m = make(map[string]int)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	file := Parse(input)
	v := 0
	for _, line := range file {
		v += solve(line, false)
	}
	return v
}

func task2(input string) int {
	file := Parse(input)
	v := 0
	for _, line := range file {
		e := expand(line, 7)
		v += solve(e, false)
	}
	return v
}

func Parse(input string) (lines []string) {
	lines = []string{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func expand(line string, iter int) string {
	s := strings.Split(line, " ")[0]
	t := strings.Split(line, " ")[1]
	s1 := s
	t1 := t
	for i := 0; i < iter-1; i++ {
		s1 += "?" + s
		t1 += "," + t
	}
	return s1 + " " + t1
}

func solve(line string, prev bool) int {
	if val, ok := m[line]; !prev && ok {
		return val
	}
	s := strings.Split(line, " ")[0]
	t := strings.Split(strings.Split(line, " ")[1], ",")
	g := []int{}
	for _, v := range t {
		tv, _ := strconv.Atoi(v)
		g = append(g, tv)
	}

	//last bit step in the gear chain
	if len(s) == 1 && len(g) > 0 {
		if s[0] == '#' {
			// current one is working gear
			if len(g) == 1 && g[0] == 1 {
				return 1
			} else {
				return 0
			}
		} else if s[0] == '.' {
			if len(g) == 1 && g[0] == 0 {
				return 1
			} else {
				return 0
			}
		}
	} else if len(g) == 0 {
		//groups have ranout
		return 0
	}
	if s[0] == '#' {
		if g[0] == 0 {
			return 0
		}
		g[0]--
		return solve(backToLine(s[1:], g), true)
	} else if s[0] == '.' {
		if prev == true {
			if g[0] == 0 {
				return solve(backToLine(s[1:], g[1:]), false)
			} else {
				return 0
			}
		}
		return solve(backToLine(s[1:], g), false)
	} else if s[0] == '?' {
		result := solve("#"+backToLine(s[1:], g), prev) + solve(backToLine("."+s[1:], g), prev)
		if !prev {
			m[line] = result
		}
		return result
	}
	return 0
}

func backToLine(s string, g []int) string {
	r := []string{}
	for _, e := range g {
		r = append(r, strconv.Itoa(e))
	}
	return s + " " + strings.Join(r, ",")

}
