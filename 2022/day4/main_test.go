package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

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
