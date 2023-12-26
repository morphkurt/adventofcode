package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

func TestTask1(t *testing.T) {
	got := task1(i, int64(7), int64(27))
	want := 2

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i, int64(7), int64(27))
	want := 47

	assert.Equal(t, want, got)
}
