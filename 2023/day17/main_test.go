package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 102

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 94

	assert.Equal(t, want, got)
}
