package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var taskOneInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

func TestTask1(t *testing.T) {
	got := task1(taskOneInput)
	want := 24
	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(taskOneInput)
	want := 93
	assert.Equal(t, want, got)
}
