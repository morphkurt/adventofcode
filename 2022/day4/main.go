package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task2(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {

		left := strings.Split(line, ",")[0]
		right := strings.Split(line, ",")[1]
		leftValues := strings.Split(left, "-")
		RightValues := strings.Split(right, "-")

		leftStart, _ := strconv.Atoi(leftValues[0])
		leftEnd, _ := strconv.Atoi(leftValues[1])

		rightStart, _ := strconv.Atoi(RightValues[0])
		rightEnd, _ := strconv.Atoi(RightValues[1])
		values := []int{leftStart, leftEnd, rightStart, rightEnd}

		leftFirstByte, leftSecondByte, rightFirstByte, rightSecondByte := uint64(0), uint64(0), uint64(0), uint64(0)

		for j := values[0]; j <= values[1]; j++ {
			if j > 64 {
				leftFirstByte = leftFirstByte | 1<<(j-65)
			} else {
				leftSecondByte = leftSecondByte | 1<<(j-1)
			}
		}

		for j := values[2]; j <= values[3]; j++ {
			if j > 64 {
				rightFirstByte = rightFirstByte | 1<<(j-65)
			} else {
				rightSecondByte = rightSecondByte | 1<<(j-1)
			}
		}

		overlapFirstByte := leftFirstByte & rightFirstByte
		overlapSecondByte := leftSecondByte & rightSecondByte

		overlapBits := bits.OnesCount(uint(overlapFirstByte)) + bits.OnesCount(uint(overlapSecondByte))
		if overlapBits > 0 {
			result++
		}
	}
	return result
}

func task1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, line := range lines {

		left := strings.Split(line, ",")[0]
		right := strings.Split(line, ",")[1]
		leftValues := strings.Split(left, "-")
		RightValues := strings.Split(right, "-")

		leftStart, _ := strconv.Atoi(leftValues[0])
		leftEnd, _ := strconv.Atoi(leftValues[1])

		rightStart, _ := strconv.Atoi(RightValues[0])
		rightEnd, _ := strconv.Atoi(RightValues[1])
		values := []int{leftStart, leftEnd, rightStart, rightEnd}

		leftFirstByte, leftSecondByte, rightFirstByte, rightSecondByte := uint64(0), uint64(0), uint64(0), uint64(0)

		for j := values[0]; j <= values[1]; j++ {
			if j > 64 {
				leftFirstByte = leftFirstByte | 1<<(j-65)
			} else {
				leftSecondByte = leftSecondByte | 1<<(j-1)
			}
		}

		for j := values[2]; j <= values[3]; j++ {
			if j > 64 {
				rightFirstByte = rightFirstByte | 1<<(j-65)
			} else {
				rightSecondByte = rightSecondByte | 1<<(j-1)
			}
		}

		overlapFirstByte := leftFirstByte & rightFirstByte
		overlapSecondByte := leftSecondByte & rightSecondByte

		overlapBits := bits.OnesCount(uint(overlapFirstByte)) + bits.OnesCount(uint(overlapSecondByte))
		leftCount := bits.OnesCount(uint(leftFirstByte)) + bits.OnesCount(uint(leftSecondByte))
		rightCount := bits.OnesCount(uint(rightFirstByte)) + bits.OnesCount(uint(rightSecondByte))
		min := math.Min(float64(leftCount), float64(rightCount))
		if overlapBits == int(min) {
			result++
		}
	}
	return result
}
