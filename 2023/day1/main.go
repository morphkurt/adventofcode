package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	task1()
	task2()
}

func task2() {

	m := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		startValue := 0
		minIndex := math.MaxInt32
		maxIndex := -1
		endValue := 0
		s := scanner.Text()
		for i, val := range m {
			index := strings.Index(s, val)
			lastIndex := strings.LastIndex(s, val)
			digitIndex := strings.Index(s, strconv.Itoa(i+1))
			lastDigitIndex := strings.LastIndex(s, strconv.Itoa(i+1))
			if index > -1 && index < minIndex {
				minIndex = index
				startValue = i + 1
			}
			if digitIndex > -1 && digitIndex < minIndex {
				minIndex = digitIndex
				startValue = i + 1
			}
			if lastIndex > -1 && lastIndex > maxIndex {
				maxIndex = lastIndex
				endValue = i + 1
			}
			if lastDigitIndex > -1 && lastDigitIndex > maxIndex {
				maxIndex = lastDigitIndex
				endValue = i + 1
			}

		}
		sum += endValue + startValue*10
	}
	fmt.Println("task two:", sum)
}

func task1() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		startValue := 0
		minIndex := math.MaxInt32
		maxIndex := -1
		endValue := 0
		s := scanner.Text()
		for i := 0; i <= 9; i++ {
			digitIndex := strings.Index(s, strconv.Itoa(i+1))
			lastIndex := strings.LastIndex(s, strconv.Itoa(i+1))
			if digitIndex > -1 && digitIndex < minIndex {
				minIndex = digitIndex
				startValue = i + 1
			}
			if lastIndex > -1 && lastIndex > maxIndex {
				maxIndex = lastIndex
				endValue = i + 1
			}

		}
		sum += (endValue) + (startValue * 10)
	}
	fmt.Println("task one:", sum)
}
