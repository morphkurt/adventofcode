package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `Time:      7  15   30
Distance:  9  40  200`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 0

	assert.Equal(t, got, want)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 0

	assert.Equal(t, got, want)
}
