package main

import (
	"fmt"
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

func task2(input string) int {
	c := strings.Split(input, "\n")
	result := 0

	for i := 0; i < len(c); i = i + 3 {
		common := []rune{}
		for _, f := range c[i] {
			for _, s := range c[i+1] {
				for _, t := range c[i+2] {
					if f == s && s == t && f == t {
						common = append(common, f)
					}
				}
			}
		}
		if len(common) > 0 {
			if common[0]-'a' > 0 {
				result += int(common[0]-'a') + 1
			} else {
				result += int(common[0]-'A') + 27
			}
		}
	}

	return result
}

func task1(input string) int {
	c := strings.Split(input, "\n")
	result := 0
	for _, v := range c {
		r := []rune(v)
		left := r[:(len(r) / 2)]
		right := r[(len(r) / 2):]
		common := []rune{}
		for _, li := range left {
			for _, ri := range right {
				if li == ri {
					common = append(common, li)
				}
			}
		}
		if len(common) > 0 {
			if common[0]-'a' > 0 {
				result += int(common[0]-'a') + 1
			} else {
				result += int(common[0]-'A') + 27
			}
		}
	}

	return result
}
