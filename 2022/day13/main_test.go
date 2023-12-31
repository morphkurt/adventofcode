package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var taskOneInput = `aabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestTask1(t *testing.T) {
	got := task1(taskOneInput)
	want := 31
	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(taskOneInput)
	want := 29
	assert.Equal(t, want, got)
}
