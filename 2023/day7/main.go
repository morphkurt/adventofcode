package main

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Hand struct {
	Stack string
	Bid   int
	Type  int
}

var StackOne = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
var StackTwo = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int64 {
	var m = StackOne
	stacks := parse(input, 1)
	sort.Slice(stacks, func(i, j int) bool {
		if stacks[i].Type != stacks[j].Type {
			return stacks[i].Type > stacks[j].Type
		}
		vj := strings.Split(stacks[j].Stack, "")
		for r, vi := range strings.Split(stacks[i].Stack, "") {
			if vi != vj[r] {
				return getCardValue(vi, m) > getCardValue(vj[r], m)
			}
		}
		return true
	})
	win := int64(0)
	for i, w := range stacks {
		win += int64(w.Bid) * (int64(i) + 1)
	}
	return win
}

func task2(input string) int64 {
	m := StackTwo
	stacks := parse(input, 2)
	sort.Slice(stacks, func(i, j int) bool {
		if stacks[i].Type != stacks[j].Type {
			return stacks[i].Type > stacks[j].Type
		}
		vj := strings.Split(stacks[j].Stack, "")
		for r, vi := range strings.Split(stacks[i].Stack, "") {
			if vi != vj[r] {
				return getCardValue(vi, m) > getCardValue(vj[r], m)
			}
		}
		return true
	})
	win := int64(0)
	for i, w := range stacks {
		win += int64(w.Bid) * (int64(i) + 1)
	}
	return win
}

func (h *Hand) ProcessStackTask1() {
	m := StackOne
	v := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cards := strings.Split(h.Stack, "")
	for _, c := range cards {
		for j, e := range m {
			if e == c {
				v[j] = v[j] + 1
			}
		}
	}
	Score(h, v)
}

func (h *Hand) ProcessStackTask2() {
	m := StackTwo
	v := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	cards := strings.Split(h.Stack, "")
	for _, c := range cards {
		for j, e := range m {
			if e == c {
				v[j] = v[j] + 1
			}
		}
	}
	mi, mv := GetMax(v[:12])
	if mv > 0 && mi != 12 {
		// add wild card value to highest card
		v[mi] = v[mi] + v[12]
		// set the wildcard to zero
		v[12] = 0
	}
	Score(h, v)
}

func parse(input string, task int) []Hand {
	hands := []Hand{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		hand := strings.Split(scanner.Text(), " ")
		c := hand[0]
		b, _ := strconv.Atoi(hand[1])
		h := Hand{
			Stack: c,
			Bid:   b,
			Type:  0,
		}
		if task == 1 {
			h.ProcessStackTask1()
		} else {
			h.ProcessStackTask2()
		}

		hands = append(hands, h)
	}
	return hands
}

func ContainsTwoPairs(c []int) int {
	numPairs := 0
	for _, e := range c {
		if e == 2 {
			numPairs++
		}
	}
	return numPairs
}

func getCardValue(k string, m []string) int {
	for i, c := range m {
		if c == k {
			return i
		}
	}
	return -1
}

func GetMax(c []int) (i, v int) {
	idx := 0
	val := math.MinInt
	for i, v := range c {
		if v > val {
			idx = i
			val = v
		}
	}
	return idx, val
}

func Score(h *Hand, v []int) {
	if slices.Contains(v, 5) {
		//fullhouse
		h.Type = 1
	} else if slices.Contains(v, 4) {
		//fourofkind
		h.Type = 2
	} else if slices.Contains(v, 3) && slices.Contains(v, 2) {
		//fullhouse
		h.Type = 3
	} else if slices.Contains(v, 3) && !slices.Contains(v, 2) {
		//three of kind
		h.Type = 4
	} else if ContainsTwoPairs(v) == 2 {
		//twopair
		h.Type = 5
	} else if ContainsTwoPairs(v) == 1 {
		//onepair
		h.Type = 6
	} else {
		//highcard
		h.Type = 7
	}
}
