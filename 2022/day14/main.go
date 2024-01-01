package main

import (
	"encoding/json"
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
	result := 0
	packets := parse(input)
	for i, p := range packets {
		if (cmp(p[0], p[1])) <= 0 {
			result += i + 1
		}

	}
	return result
}

func task2(input string) int {
	result := 1
	packets := parse(input)
	allPackets := []any{}
	for _, p := range packets {
		allPackets = append(allPackets, p...)
	}
	allPackets = append(allPackets, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(allPackets, func(i, j int) bool { return cmp(allPackets[i], allPackets[j]) < 0 })
	for i, p := range allPackets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			result *= i + 1
		}
	}
	return result
}

func parse(input string) [][]any {
	out := [][]any{}
	blocks := strings.Split(input, "\n\n")
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		left := lines[0]
		right := lines[1]
		var leftItem any
		var rightItem any
		json.Unmarshal([]byte(left), &leftItem)
		json.Unmarshal([]byte(right), &rightItem)
		out = append(out, []any{leftItem, rightItem})
	}
	return out
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}
