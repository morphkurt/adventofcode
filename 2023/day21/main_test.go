package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func TestTask1(t *testing.T) {
	got := task1(i, 6)
	want := 16

	assert.Equal(t, want, got)
}
