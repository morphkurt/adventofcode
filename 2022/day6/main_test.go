package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 7

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 19

	assert.Equal(t, want, got)
}
