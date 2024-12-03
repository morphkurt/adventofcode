package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 48

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 4

	assert.Equal(t, want, got)
}
