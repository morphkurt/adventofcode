package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

var i2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 2

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got, _ := task2(i2)
	want := 6

	assert.Equal(t, want, got)
}

func TestTask2Yolo(t *testing.T) {

	got, _ := task2BruteForce(i2)
	want := 6

	assert.Equal(t, want, got)
}

func TestFindGcd(t *testing.T) {

	got := FindGcd(21883, 13019)
	want := 277

	assert.Equal(t, want, got)
}

func TestParsing(t *testing.T) {
	Directions, Nodes := ParseInput(i)
	ExpectedDirections := []string{"R", "L"}
	ExpectedNodes := []node{
		{
			Name:      "AAA",
			LeftNode:  "BBB",
			RightNode: "CCC",
		},
		{
			Name:      "BBB",
			LeftNode:  "DDD",
			RightNode: "EEE",
		},
		{
			Name:      "CCC",
			LeftNode:  "ZZZ",
			RightNode: "GGG",
		},
		{
			Name:      "DDD",
			LeftNode:  "DDD",
			RightNode: "DDD",
		},
		{
			Name:      "EEE",
			LeftNode:  "EEE",
			RightNode: "EEE",
		},
		{
			Name:      "GGG",
			LeftNode:  "GGG",
			RightNode: "GGG",
		},
		{
			Name:      "ZZZ",
			LeftNode:  "ZZZ",
			RightNode: "ZZZ",
		},
	}
	assert.Equal(t, Directions, ExpectedDirections)
	assert.Equal(t, Nodes, ExpectedNodes)
}
