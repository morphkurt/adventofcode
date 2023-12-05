package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(?s)Card[ ]+(\d+): ([\w ]+) | ([\w ]+)`)
var digitRegex = regexp.MustCompile(`(?m)(\d+)`)

func main() {
	task1()
	task2()
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
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		l := scanner.Text()
		won := []int{}
		match := re.FindAllStringSubmatch(l, -1)
		//cardNum, _ := strconv.Atoi(match[0][1])
		winningsNumbers := digitRegex.FindAllString(match[0][2], -1)
		drawnNumbers := digitRegex.FindAllString(match[1][0], -1)
		for _, d := range drawnNumbers {
			for _, w := range winningsNumbers {
				if d == w {
					val, _ := strconv.Atoi(d)
					won = append(won, val)
				}
			}
		}
		sum += getPoints(len(won))
	}
	fmt.Printf("task 1: %d\n", sum)
}

func task2() {

	cards := []string{}
	uniqueCards := []int{}

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

	for scanner.Scan() { // internally, it advances token based on sperator
		l := scanner.Text()
		cards = append(cards, l)
		uniqueCards = append(uniqueCards, 1)
	}
	for i, c := range cards {
		iteration := uniqueCards[i]
		for j := 0; j < iteration; j++ {
			won := []int{}
			match := re.FindAllStringSubmatch(c, -1)
			//cardNum, _ := strconv.Atoi(match[0][1])
			winningsNumbers := digitRegex.FindAllString(match[0][2], -1)
			drawnNumbers := digitRegex.FindAllString(match[1][0], -1)
			for _, d := range drawnNumbers {
				for _, w := range winningsNumbers {
					if d == w {
						val, _ := strconv.Atoi(d)
						won = append(won, val)
					}
				}
			}
			for o, _ := range won {
				if i+o+1 < len(uniqueCards) {
					uniqueCards[i+o+1] += 1
				}
			}
		}

	}
	sum := 0
	for _, c := range uniqueCards {
		sum += c
	}
	fmt.Printf("task 2 :%d\n", sum)
}

func getPoints(n int) int {
	if n == 0 {
		return 0
	}
	p := 0
	for n >= 0 {
		p += int(math.Pow(2, float64(n-2)))
		n--
	}
	return p + 1
}
