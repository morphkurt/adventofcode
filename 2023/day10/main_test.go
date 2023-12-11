package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var p2 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 8

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(p2)
	want := 10

	assert.Equal(t, want, got)
}

func TestFindOrigin(t *testing.T) {
	m, _ := ReadMap(i)
	x, y := FindOrigin(m)
	got := []int{x, y}
	want := []int{0, 2}

	assert.Equal(t, want, got)
}

func TestGetNextLocation(t *testing.T) {
	m, _ := ReadMap(i)
	x, y := FindOrigin(m)
	got := GetNextLocation(m, x, y)
	want := []Location{
		{
			x:    0,
			y:    3,
			Type: "|",
		},
		{
			x:    1,
			y:    2,
			Type: "J",
		},
	}
	assert.Equal(t, want, got)
}
