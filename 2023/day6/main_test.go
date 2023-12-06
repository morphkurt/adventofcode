package main

import (
	"reflect"
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

	if !reflect.DeepEqual(times, expectedTimes) {
		t.Errorf("got %d, wanted %d", times, expectedTimes)
	}

	if !reflect.DeepEqual(distances, expectedDistances) {
		t.Errorf("got %d, wanted %d", distances, expectedDistances)
	}
}
