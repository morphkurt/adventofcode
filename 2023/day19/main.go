package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

const (
	X        = iota
	M        = iota
	A        = iota
	S        = iota
	ACCEPT   = iota
	REJECT   = iota
	WORKFLOW = iota
	GT       = iota
	LT       = iota
)

type WorkFlow struct {
	name  string
	rules []Rule
}

type Rule struct {
	catagory int
	logic    int
	value    int
	nextflow string
}

type Part struct {
	attributes map[int]int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	result := 0
	flows, parts := parse(input)
	for _, p := range parts {
		ok := p.Process(flows)
		if ok {
			result += p.attributes[X] + p.attributes[M] + p.attributes[A] + p.attributes[S]
		}
	}
	return result
}

func task2(input string) int {
	flows, _ := parse(input)
	min := map[int][]int{}
	max := map[int][]int{}

	for _, rules := range flows {
		for _, rule := range rules {
			if rule.logic == LT {
				min[rule.catagory] = append(min[rule.catagory], rule.value)
				max[rule.catagory] = append(max[rule.catagory], rule.value-1)
			}
			if rule.logic == GT {
				max[rule.catagory] = append(max[rule.catagory], rule.value)
				min[rule.catagory] = append(min[rule.catagory], rule.value+1)
			}
		}
	}

	for p := range min {
		min[p] = append(min[p], 1)
		sort.Slice(min[p], func(i, j int) bool {
			return min[p][i] < min[p][j]
		})
	}
	for p := range max {
		max[p] = append(max[p], 4000)
		sort.Slice(max[p], func(i, j int) bool {
			return max[p][i] < max[p][j]
		})
	}

	n := 0
	for xi, x := range min[X] {
		for mi, m := range min[M] {
			for si, s := range min[S] {
				for ai, a := range min[A] {
					attr := map[int]int{X: x, M: m, S: s, A: a}
					part := Part{attributes: attr}
					if part.Process(flows) {
						n += (max[A][ai] - a + 1) *
							(max[S][si] - s + 1) *
							(max[M][mi] - m + 1) *
							(max[X][xi] - x + 1)
					}
				}
			}
		}
	}

	return n
}

func (part Part) Process(flows map[string][]Rule) bool {
	inFlow := flows["in"]
	nextFlow := inFlow
	resultFound := false
	for !resultFound {
		for _, flow := range nextFlow {
			flow, ok := flow.Evaluate(part)
			if flow == "A" {
				return true
			}
			if flow == "R" {
				return false
			}
			if ok {
				nextFlow = flows[flow]
				break
			}
		}
	}
	return false
}

func (rule Rule) Evaluate(part Part) (string, bool) {
	cat := rule.catagory
	op := rule.logic
	val := rule.value
	nextflow := rule.nextflow
	switch op {
	case GT:
		if part.attributes[cat] > val {
			return nextflow, true
		}
	case LT:
		if part.attributes[cat] < val {
			return nextflow, true
		}
	default:
		return nextflow, true
	}
	return "", false
}

func parse(input string) (map[string][]Rule, []Part) {
	re := regexp.MustCompile(`(?m){(\w+)=(\d+),(\w+)=(\d+),(\w+)=(\d+),(\w+)=(\d+)}`)
	flows := make(map[string][]Rule)
	parts := []Part{}
	rules := strings.Split(input, "\n\n")[0]
	items := strings.Split(input, "\n\n")[1]

	for _, item := range strings.Split(items, "\n") {
		attributes := make(map[int]int)
		match := re.FindAllStringSubmatch(item, -1)
		x, _ := strconv.Atoi(match[0][2])
		m, _ := strconv.Atoi(match[0][4])
		a, _ := strconv.Atoi(match[0][6])
		s, _ := strconv.Atoi(match[0][8])
		attributes[X] = x
		attributes[M] = m
		attributes[A] = a
		attributes[S] = s
		parts = append(parts, Part{
			attributes: attributes,
		})
	}
	ruleLines := strings.Split(rules, "\n")
	for _, line := range ruleLines {
		split := strings.Split(line, "{")
		name := split[0]
		split[1] = strings.ReplaceAll(split[1], "}", "")
		rules := strings.Split(split[1], ",")
		ruleArray := []Rule{}
		for _, r := range rules {
			ri := Rule{}
			if strings.Contains(r, ":") {
				i := strings.Split(r, ":")
				itemType := strings.FieldsFunc(i[0], Split)[0]
				value := strings.FieldsFunc(i[0], Split)[1]
				valueInt, _ := strconv.Atoi(value)
				cat := 0
				switch itemType {
				case "a":
					cat = A
				case "x":
					cat = X
				case "s":
					cat = S
				case "m":
					cat = M
				}
				operation := rune(i[0][1])
				logic := 0
				switch operation {
				case '>':
					logic = GT
				case '<':
					logic = LT
				}
				outFlow := i[1]
				ri.catagory = cat
				ri.logic = logic
				ri.nextflow = outFlow
				ri.value = valueInt

			} else {
				ri.nextflow = r
			}
			ruleArray = append(ruleArray, ri)

		}
		flows[name] = ruleArray
	}
	return flows, parts
}

func Split(r rune) bool {
	return r == '<' || r == '>'
}
