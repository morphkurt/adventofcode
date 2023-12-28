package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `30373
25512
65332
33549
35390`

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
