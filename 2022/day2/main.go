package main

import (
	"fmt"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

const (
	ROCK    = iota
	PAPER   = iota
	SCISSOR = iota
)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task2(input string) int {
	// A B C
	// B C A //win
	// C A B //loose
	// A B C //draw
	loose := []int{'C', 'A', 'B'}
	win := []int{'B', 'C', 'A'}
	draw := []int{'A', 'B', 'C'}

	hands := Parse(input)
	result := 0
	myhand := 0
	for _, v := range hands {
		end := v[1]
		switch end {
		case "X": //loose
			myhand = loose[int(v[0][0]-'A')]
			result += 0
		case "Y": //draw
			myhand = draw[int(v[0][0]-'A')]
			result += 3
		case "Z": //win
			myhand = win[int(v[0][0]-'A')]
			result += 6
		}
		result += (myhand - 'A') + 1
	}
	return result
}

func task1(input string) int {
	hands := Parse(input)
	result := 0
	for _, v := range hands {
		if DidWin(v) == 'W' {
			result += 6
		}
		if DidWin(v) == 'D' {
			result += 3
		}
		switch v[1] {
		case "X":
			result += 1
		case "Y":
			result += 2
		case "Z":
			result += 3
		}
	}
	return result
}

func DidWin(hand []string) rune {
	if hand[0][0]-'A' == hand[1][0]-'X' {
		return 'D'
	}
	switch hand[0] {
	case "B": //paper
		if hand[1] == "Z" {
			return 'W'
		}
	case "A": //rock
		if hand[1] == "Y" {
			return 'W'
		}
	case "C": //scissor
		if hand[1] == "X" {
			return 'W'
		}
	}
	return 'L'
}
func Parse(input string) [][]string {
	out := [][]string{}
	game := strings.Split(input, "\n")
	for _, v := range game {
		elv := strings.Split(v, " ")[0]
		my := strings.Split(v, " ")[1]
		out = append(out, []string{elv, my})
	}
	return out
}
