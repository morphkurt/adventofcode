package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `    [D]        
[N] [C]    
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := "CMZ"

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := "MCD"

	assert.Equal(t, want, got)
}
