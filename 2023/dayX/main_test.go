package main

import (
	"testing"
)

var i = ``

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 0
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
