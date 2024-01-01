package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var taskOneInput = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

func TestTask1(t *testing.T) {
	got := task1(taskOneInput)
	want := 13
	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(taskOneInput)
	want := 140
	assert.Equal(t, want, got)
}
