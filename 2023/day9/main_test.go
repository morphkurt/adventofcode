package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 114

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	want := task2(i)
	got := 2

	assert.Equal(t, got, want)
}

func TestParsing(t *testing.T) {
	want := [][]int([][]int{[]int{0, 3, 6, 9, 12, 15}, []int{1, 3, 6, 10, 15, 21}, []int{10, 13, 16, 21, 30, 45}})
	got := ParseInput(i)
	assert.Equal(t, want, got)
}
