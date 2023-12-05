package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

type digitLoc struct {
	value int
	row   int
	col   int
}

type enginePart struct {
	number int
	digit  []digitLoc
}

func main() {
	task1()
	task2()
}

func task2() {

	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	sum := 0
	br := bufio.NewReader(f)
	var digits []digitLoc
	var matrix [][]bool
	var engineParts []enginePart
	row := 0
	col := 0
	var matrixRow []bool
	for {

		b, err := br.ReadByte()
		if !isReturn(int(b)) {
			if isStar(int(b)) {
				matrixRow = append(matrixRow, true)
			} else {
				matrixRow = append(matrixRow, false)
			}
		}

		if isReturn(int(b)) {
			row++
			col = -1
			matrix = append(matrix, matrixRow)
			matrixRow = []bool{}
		}

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if isDigit(int(b)) {
			var digit digitLoc
			digit.value = int(b) - 48
			digit.row = row
			digit.col = col
			digits = append(digits, digit)
		} else {
			value := 0
			for i, d := range digits {
				value += d.value * int(math.Pow10(len(digits)-i-1))
			}
			if len(digits) > 0 {
				var part enginePart
				part.digit = digits
				part.number = value
				engineParts = append(engineParts, part)
			}
			digits = []digitLoc{}
		}

		// process the one byte b
		col += 1
		if err != nil {
			// end of file
			break
		}
	}

	for x, c := range matrix {
		for y, r := range c {
			gears := []int{}
			for _, p := range engineParts {
				validGear := false
				for _, d := range p.digit {
					if r {
						if isNeighbour(matrix, x, y, d.row, d.col) {
							validGear = true
						}
					}
				}
				if validGear {
					gears = append(gears, p.number)

				}
				validGear = false
			}
			if len(gears) == 2 {
				sum += gears[0] * gears[1]
			}
			gears = []int{}
		}
	}

	fmt.Println("task 2,", sum)
}

func task1() {

	f, err := os.Open("input")

	if err != nil {
		panic(err)
	}
	sum := 0
	br := bufio.NewReader(f)
	var digits []digitLoc
	var matrix [][]bool
	var engineParts []enginePart
	row := 0
	col := 0
	var matrixRow []bool
	for {

		b, err := br.ReadByte()
		if !isReturn(int(b)) {
			if isSymbol(int(b)) {
				matrixRow = append(matrixRow, true)
			} else {
				matrixRow = append(matrixRow, false)
			}
		}

		if isReturn(int(b)) {
			row++
			col = -1
			matrix = append(matrix, matrixRow)
			matrixRow = []bool{}
		}

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if isDigit(int(b)) {
			var digit digitLoc
			digit.value = int(b) - 48
			digit.row = row
			digit.col = col
			digits = append(digits, digit)
		} else {
			value := 0
			for i, d := range digits {
				value += d.value * int(math.Pow10(len(digits)-i-1))
			}
			if len(digits) > 0 {
				var part enginePart
				part.digit = digits
				part.number = value
				engineParts = append(engineParts, part)
			}
			digits = []digitLoc{}
		}

		// process the one byte b
		col += 1
		if err != nil {
			// end of file
			break
		}
	}

	for _, p := range engineParts {
		validPart := false
		for _, d := range p.digit {
			if isNeighbourSymbol(matrix, d.row, d.col) {
				validPart = true
			}
		}
		if validPart {
			sum += p.number
		}
		validPart = false
	}
	fmt.Println("task 1,", sum)
}

func isDigit(b int) bool {
	if b > 47 && b < 58 {
		return true
	}
	return false
}

func isSymbol(b int) bool {
	if isDigit(b) || b == 46 || isReturn(b) {
		return false
	}
	return true
}

func isStar(b int) bool {
	if b == 42 {
		return true
	}
	return false
}

func isReturn(b int) bool {
	if b == 10 {
		return true
	}
	return false
}

func isNeighbourSymbol(matrix [][]bool, row int, col int) bool {
	n := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, +1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	length := len(matrix[0][:])
	height := len(matrix[:][0])
	for _, c := range n {
		x := row + c[0]
		y := col + c[1]
		if x >= 0 && y >= 0 && x < length-1 && y < height-1 && matrix[x][y] {
			return true
		}
	}
	return false
}

func isNeighbour(matrix [][]bool, row int, col int, digCol int, digRow int) bool {
	n := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, +1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, c := range n {
		x := row + c[0]
		y := col + c[1]
		if x == digCol && y == digRow {
			return true
		}
	}
	return false
}
