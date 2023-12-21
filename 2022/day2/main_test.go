package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `A Y
B X
C Z`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 15

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 12

	assert.Equal(t, want, got)
}
