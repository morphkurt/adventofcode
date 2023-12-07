package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := int64(6440)

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {

	got := task2(i)
	want := int64(5905)

	assert.Equal(t, got, want)
}

func TestParsing(t *testing.T) {
	GotHands := parse(i, 1)

	ExpectedHand := []Hand{
		Hand{
			Stack: "32T3K",
			Bid:   765,
			Type:  6,
		},
		Hand{
			Stack: "T55J5",
			Bid:   684,
			Type:  4,
		},
		Hand{
			Stack: "KK677",
			Bid:   28,
			Type:  5,
		},
		Hand{
			Stack: "KTJJT",
			Bid:   220,
			Type:  5,
		},
		Hand{
			Stack: "QQQJA",
			Bid:   483,
			Type:  4,
		},
	}

	assert.Equal(t, GotHands, ExpectedHand)
}
