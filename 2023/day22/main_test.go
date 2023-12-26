package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `0,0,2~2,0,2
0,1,6~2,1,6
1,1,8~1,1,9
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
1,0,1~1,2,1`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 5

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	//got := task1(i)
	//want := 16

	//assert.Equal(t, want, got)
}

func TestDoesOverlap(t *testing.T) {
	r1x1, r1x2, r1y1, r1y2 := 0, 1, 0, 3
	r2x1, r2x2, r2y1, r2y2 := 0, 3, 2, 3

	got := DoesOverlap(r1x1, r1x2, r2x1, r2x2, r1y1, r1y2, r2y1, r2y2)
	want := true
	assert.Equal(t, want, got)
}
