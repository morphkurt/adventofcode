package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type monkey struct {
	id                int
	inspections       int
	items             []uint
	operation         func(uint, uint) uint
	operationOldValue bool
	operationVal      int
	testValue         int
	success           *monkey
	fail              *monkey
}

func add(a uint, b uint) uint {
	return a + b
}

func multiply(a uint, b uint) uint {
	return a * b
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	monkeys := parse(input)
	inspections := []int{}
	for i := 1; i <= 20; i++ {
		processPart1(monkeys, 3)
	}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func task2(input string) int {
	monkeys := parse(input)
	inspections := []int{}
	mod := 1
	for _, m := range monkeys {
		mod *= m.testValue
	}
	for i := 1; i <= 10000; i++ {
		processPart2(monkeys, mod)
	}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func processPart1(monkeys []*monkey, worryReduction int) {
	for i := range monkeys {
		for len(monkeys[i].items) > 0 {
			item := monkeys[i].items[0]
			monkeys[i].items = monkeys[i].items[1:]
			val := uint(monkeys[i].operationVal)
			if monkeys[i].operationOldValue {
				val = item
			}
			worryLevel := monkeys[i].operation(item, val) / uint(worryReduction)
			if worryLevel%uint(monkeys[i].testValue) == 0 {
				monkeys[i].success.items = append(monkeys[i].success.items, worryLevel)
			} else {
				monkeys[i].fail.items = append(monkeys[i].fail.items, worryLevel)
			}
			monkeys[i].inspections++
		}
	}
}

func processPart2(monkeys []*monkey, worryReduction int) {
	for i := range monkeys {
		for len(monkeys[i].items) > 0 {
			item := monkeys[i].items[0]
			monkeys[i].items = monkeys[i].items[1:]
			val := uint(monkeys[i].operationVal)
			if monkeys[i].operationOldValue {
				val = item
			}
			worryLevel := monkeys[i].operation(item, val) % uint(worryReduction)
			if worryLevel%uint(monkeys[i].testValue) == 0 {
				monkeys[i].success.items = append(monkeys[i].success.items, worryLevel)
			} else {
				monkeys[i].fail.items = append(monkeys[i].fail.items, worryLevel)
			}
			monkeys[i].inspections++
		}
	}
}

func parse(input string) []*monkey {
	out := []*monkey{}
	monkeyMap := map[int]*monkey{}
	monkeyBlocks := strings.Split(input, "\n\n")
	for i, block := range monkeyBlocks {
		lines := strings.Split(block, "\n")
		for j, line := range lines {
			lines[j] = strings.TrimSpace(line)
		}
		monkeyBlocks[i] = strings.Join(lines, "\n")
		monkeyMap[i] = &monkey{id: i, items: []uint{}}
	}
	for _, block := range monkeyBlocks {
		lines := strings.Split(block, "\n")
		var id int
		fmt.Sscanf(lines[0], "Monkey %d:", &id)
		monkey := monkeyMap[id]
		itemsStr := strings.Split(lines[1], ":")[1]
		items := util.ToUint(strings.Split(itemsStr, ","))
		monkey.items = append(monkey.items, items...)
		operationStr := strings.ReplaceAll(lines[2], "Operation: new = old ", "")
		switch operationStr[0] {
		case '*':
			if strings.Contains(operationStr, "old") {
				monkey.operationOldValue = true
			} else {
				splitted := strings.Split(lines[2], " ")
				val := util.ToInt([]string{splitted[len(splitted)-1]})
				monkey.operationVal = val[0]
				monkey.operationOldValue = false
			}
			monkey.operation = multiply
		case '+':
			if strings.Contains(operationStr, "old") {
				monkey.operationOldValue = true
			} else {
				splitted := strings.Split(lines[2], " ")
				val := util.ToInt([]string{splitted[len(splitted)-1]})
				monkey.operationVal = val[0]
				monkey.operationOldValue = false
			}
			monkey.operation = add
		}
		var testValue int
		fmt.Sscanf(lines[3], "Test: divisible by %d", &testValue)
		monkey.testValue = testValue
		var succesMonkeyId int
		fmt.Sscanf(lines[4], "If true: throw to monkey %d", &succesMonkeyId)
		monkey.success = monkeyMap[succesMonkeyId]
		var failMonkeyId int
		fmt.Sscanf(lines[5], "If false: throw to monkey %d", &failMonkeyId)
		monkey.fail = monkeyMap[failMonkeyId]
		out = append(out, monkey)
	}
	return out
}
