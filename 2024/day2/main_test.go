package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 2

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 4

	assert.Equal(t, want, got)
}
