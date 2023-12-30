package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var cycleMap = map[string]int{
	"noop": 1,
	"addx": 2,
}

type instruction struct {
	name   string
	value  int
	cycles int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:\n%s", task2)
}

func task1(input string) int {
	cycles := []int{20, 60, 100, 140, 180, 220}
	result := 0
	instructions := parse(input)

	for _, c := range cycles {
		result += c * process(c, instructions)
	}

	return result
}

func task2(input string) string {
	instructions := parse(input)
	result := ""
	line := []rune{}
	for i := 1; i <= 240; i++ {
		r := process(i, instructions)
		if i%40 >= r && i%40 < r+3 {
			line = append(line, '#')
		} else {
			line = append(line, '.')
		}
		if (len(line) % 40) == 0 {
			result += string(line) + "\n"
			line = []rune{}
		}
	}
	return result
}

func process(cycles int, instructions []instruction) int {
	inst := instructions[0]
	instIdx := 1 //next instruction
	instCycle := inst.cycles
	reg := 1
	for i := 1; i <= cycles; i++ {
		if instCycle == 0 {
			switch inst.name {
			case "addx":
				reg += inst.value
			}
			inst = instructions[instIdx]
			instCycle = inst.cycles
			instIdx++
		}
		instCycle--
	}
	return reg
}

func parse(input string) []instruction {
	out := []instruction{}
	rows := strings.Split(input, "\n")
	for _, line := range rows {
		splitted := strings.Split(line, " ")
		var value int
		if len(splitted) > 1 {
			value, _ = strconv.Atoi(splitted[1])
		}
		out = append(out, instruction{name: splitted[0], value: value, cycles: cycleMap[splitted[0]]})
	}
	return out
}
