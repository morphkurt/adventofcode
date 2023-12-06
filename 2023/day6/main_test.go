package main

import (
	"testing"
)

var i = `Time:      7  15   30
Distance:  9  40  200`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 288

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := 71503
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestParsing(t *testing.T) {
	times, distances := parse(i)
	expectedTimes := []int{7, 15, 30}
	expectedDistances := []int{9, 40, 200}
	for i, et := range expectedTimes {
		if times[i] != et {
			t.Errorf("got %d, wanted %d", times[i], et)
		}
	}
	for i, ed := range expectedDistances {
		if distances[i] != ed {
			t.Errorf("got %d, wanted %d", distances[i], ed)
		}
	}
}
