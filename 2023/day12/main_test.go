package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 21

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 525152

	assert.Equal(t, want, got)
}
