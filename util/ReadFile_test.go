package util

import (
	"testing"
)

var inputFile = `seeds: 79 14 55 13
seeds: 79 14 55 13`

func TestReadFile(t *testing.T) {

	got := ReadFile("test_data")
	want := inputFile

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
