package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

var instMap = map[string]inst{
	"noop": {
		instFn: noop,
		cycles: 1,
	},
	"addx": {
		instFn: addx,
		cycles: 2,
	},
}

type inst struct {
	instFn func(int, int) int
	value  int
	cycles int
}

func addx(r, x int) int {
	return r + x
}

func noop(r, x int) int {
	return r
}

type cpu struct {
	instList []inst
	reg      int
	instIdx  int
	insCycle int
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
	instructions := parse(input)
	result := 0
	cpu := cpu{instList: instructions, instIdx: 0, insCycle: instructions[0].cycles, reg: 1}
	for i := 1; i <= 220; i++ {
		cpu.tick()
		if slices.Contains(cycles, i) {
			result += cpu.reg * i
		}
	}
	return result
}

func task2(input string) string {
	instructions := parse(input)
	cpu := cpu{instList: instructions, instIdx: 0, insCycle: instructions[0].cycles, reg: 1}
	result := ""
	line := []rune{}
	for i := 1; i <= 240; i++ {
		cpu.tick()
		r := cpu.reg
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

func (c *cpu) tick() {
	if c.insCycle == 0 {
		if c.instIdx < len(c.instList) {
			f := c.instList[c.instIdx].instFn
			c.reg = f(c.reg, c.instList[c.instIdx].value)
			if c.instIdx < len(c.instList)-1 {
				c.instIdx++
			}
			c.insCycle = c.instList[c.instIdx].cycles
		}
	}
	if c.insCycle > 0 {
		c.insCycle--
	}

}

func parse(input string) []inst {
	out := []inst{}
	rows := strings.Split(input, "\n")
	for _, line := range rows {
		splitted := strings.Split(line, " ")
		var value int
		if len(splitted) > 1 {
			value, _ = strconv.Atoi(splitted[1])
		}
		i := instMap[splitted[0]]
		i.value = value
		out = append(out, i)
	}
	return out
}
