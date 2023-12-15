package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type lens struct {
	name  string
	value int
}

type box struct {
	number int
	lenses []lens
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	r := 0
	steps := Parse(input)
	for _, s := range steps {
		r += Hash(s)
	}
	return r
}

func task2(input string) int {
	var boxes = []box{}
	r := 0
	for i := 0; i < 256; i++ {
		b := box{
			number: i,
			lenses: []lens{},
		}
		boxes = append(boxes, b)
	}
	steps := Parse(input)
	for _, s := range steps {
		if strings.HasSuffix(s, "-") {
			name := strings.Split(s, "-")[0]
			boxName := Hash(name)
			removeLens(boxName, name, boxes)

		} else if strings.Contains(s, "=") {
			name := strings.Split(s, "=")[0]
			focalLength, _ := strconv.Atoi(strings.Split(s, "=")[1])
			boxName := Hash(name)
			addLens(boxName, name, focalLength, boxes)
		}
	}
	for _, b := range boxes {
		boxWeight := b.number + 1
		for i, l := range b.lenses {
			r += boxWeight * l.value * (i + 1)
		}
	}
	return r
}

func Parse(input string) (steps []string) {
	input = strings.ReplaceAll(input, "\r", "")
	return strings.Split(input, ",")
}

func Hash(input string) int {
	v := 0
	for _, r := range input {
		v += int(r)
		v = (v * 17 % 256)
	}
	return v
}

func removeLens(n int, l string, boxes []box) {
	for i, v := range boxes[n].lenses {
		if v.name == l {
			boxes[n].lenses = RemoveIndex(boxes[n].lenses, i)
		}
	}
}

func addLens(n int, l string, focalLength int, boxes []box) {
	found := false
	for i, v := range boxes[n].lenses {
		if v.name == l {
			boxes[n].lenses[i].value = focalLength
			found = true
		}
	}
	if !found {
		boxes[n].lenses = append(boxes[n].lenses, lens{name: l, value: focalLength})
	}
}

func RemoveIndex(s []lens, index int) []lens {
	return append(s[:index], s[index+1:]...)
}
