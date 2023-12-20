package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
output ->`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 11687500

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 0

	assert.Equal(t, want, got)
}
