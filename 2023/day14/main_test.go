package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 136

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 64

	assert.Equal(t, want, got)
}
