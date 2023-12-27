package main

import (
	"fmt"
	"math/bits"

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
	ch := []rune(input)
	for i := 0; i < len(ch)-1; i++ {
		chTest := 0
		a := []rune{ch[i], ch[i+1], ch[i+2], ch[i+3]}
		for _, c := range a {
			chTest = chTest | 1<<int(c-'a')
		}
		if bits.OnesCount(uint(chTest)) == 4 {
			return i + 4
		}
	}
	return 0
}

func task2(input string) int {
	ch := []rune(input)
	for i := 0; i < len(ch)-14; i++ {
		chTest := 0
		a := []rune{}
		for j := 0; j < 14; j++ {
			a = append(a, ch[i+j])
		}
		for _, c := range a {
			chTest = chTest | 1<<int(c-'a')
		}
		if bits.OnesCount(uint(chTest)) == 14 {
			return i + 14
		}
	}
	return 0
}
