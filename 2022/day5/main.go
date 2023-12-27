package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Stack struct {
	id    int
	items []rune
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%s\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%s\n", task2)
}

func Parse(input string) (map[int]Stack, [][]int) {
	stack := strings.Split(input, "\n\n")[0]
	stackLine := strings.Split(stack, "\n")
	chatAt := []int{}
	stackMap := map[int]Stack{}
	for i := len(stackLine) - 1; i >= 0; i-- {
		if i == len(stackLine)-1 {
			line := []rune(stackLine[i])
			for j := 0; j < len(line); j++ {
				if line[j] >= '0' && line[j] <= '9' {
					chatAt = append(chatAt, j)
					stackMap[int(line[j]-'0')] = Stack{id: int(line[j] - '0'), items: []rune{}}
				}

			}
		} else {
			for _, k := range chatAt {
				if stackLine[i][k] >= 'A' && stackLine[i][k] <= 'Z' {
					stackItem := stackMap[int(stackLine[len(stackLine)-1][k]-'0')]
					stackItem.items = append(stackItem.items, rune(stackLine[i][k]))
					stackMap[int(stackLine[len(stackLine)-1][k]-'0')] = stackItem
				}
			}
		}

	}
	inst := strings.Split(input, "\n\n")[1]
	re := regexp.MustCompile(`(?m)move (\d+) from (\d+) to (\d+)`)
	out := [][]int{}
	for _, v := range strings.Split(inst, "\n") {
		match := re.FindAllStringSubmatch(v, -1)
		s := []string{match[0][1], match[0][2], match[0][3]}
		out = append(out, util.ToInt(s))
	}
	return stackMap, out
}

func task1(input string) string {
	m, inst := Parse(input)
	for _, v := range inst {
		b := v[0]
		s := v[1]
		d := v[2]
		source := m[s]
		destination := m[d]
		for i := 0; i < b; i++ {
			a := source.items
			destination.items = append(destination.items, a[len(a)-1])
			a = a[:len(a)-1]
			m[d] = destination
			source.items = a
			m[s] = source
		}
	}
	out := []rune{}
	for i := 1; i <= len(m); i++ {
		out = append(out, m[i].items[len(m[i].items)-1])
	}

	return string(out)
}

func task2(input string) string {
	m, inst := Parse(input)
	for _, v := range inst {
		b := v[0]
		s := v[1]
		d := v[2]
		source := m[s]
		destination := m[d]
		split := source.items[len(source.items)-b : len(source.items)]
		destination.items = append(destination.items, split...)
		source.items = source.items[:len(source.items)-b]
		m[s] = source
		m[d] = destination
	}
	out := []rune{}
	for i := 1; i <= len(m); i++ {
		out = append(out, m[i].items[len(m[i].items)-1])
	}
	return string(out)
}
