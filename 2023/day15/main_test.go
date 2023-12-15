package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 1320

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 145

	assert.Equal(t, want, got)
}

func TestHash(t *testing.T) {

	got := Hash("HASH")
	want := 52

	assert.Equal(t, want, got)
}
