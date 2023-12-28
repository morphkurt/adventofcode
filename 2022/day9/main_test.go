package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 21
	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 8

	assert.Equal(t, want, got)
}
