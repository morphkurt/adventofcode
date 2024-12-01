package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 11

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 31

	assert.Equal(t, want, got)
}
