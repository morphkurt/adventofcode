package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = ``

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 0

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 0

	assert.Equal(t, got, want)
}
