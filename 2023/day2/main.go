package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	task1()
	task2()
}

func task2() {

	m := make(map[string]int)

	m["red"] = 12
	m["green"] = 13
	m["blue"] = 14

	minColor := make(map[string]int)

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
	for scanner.Scan() { // internally, it advances token based on sperator
		s := strings.Split(scanner.Text(), ":")
		set := strings.Split(s[1], ";")
		minColor["red"] = 0
		minColor["green"] = 0
		minColor["blue"] = 0
		for _, s := range set {
			dice := strings.Split(s, ",")
			for _, d := range dice {
				color := strings.Split(d, " ")
				n, _ := strconv.Atoi(color[1])
				colorValue := color[2]
				if n > minColor[colorValue] {
					minColor[colorValue] = n
				}
			}
		}
		sum += minColor["red"] * minColor["green"] * minColor["blue"]
	}
	fmt.Println("task2:", sum)
}

func task1() {

	m := make(map[string]int)

	m["red"] = 12
	m["green"] = 13
	m["blue"] = 14

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
	for scanner.Scan() { // internally, it advances token based on sperator
		s := strings.Split(scanner.Text(), ":")
		game, _ := strconv.Atoi(strings.Replace(s[0], "Game ", "", 1))
		set := strings.Split(s[1], ";")
		invalid := false
		for _, s := range set {
			dice := strings.Split(s, ",")
			for _, d := range dice {
				color := strings.Split(d, " ")
				n, _ := strconv.Atoi(color[1])
				colorValue := color[2]
				if n > m[colorValue] {
					invalid = true
				}
			}
		}
		if !invalid {
			sum += game
		}
	}
	fmt.Println("task1:", sum)
}
