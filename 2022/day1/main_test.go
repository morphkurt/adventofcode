package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 24000

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 45000

	assert.Equal(t, want, got)
}
