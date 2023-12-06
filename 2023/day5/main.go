package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type mapper struct {
	dest   []int64
	source []int64
	length []int64
}

func main() {
	input := util.ReadFile("input")
	val1 := task1(input)
	fmt.Println(val1)
	val := task2(input)
	fmt.Println(val)
}

func task2(input string) int64 {
	h := make(map[string]mapper)
	scanner := bufio.NewScanner(strings.NewReader(input))
	expectMap := false
	dest := []int64{}
	source := []int64{}
	length := []int64{}
	seeds := []int64{}
	mapName := ""

	for scanner.Scan() { // internally, it advances token based on sperator
		line := scanner.Text()
		if strings.Contains(line, "seeds:") && !expectMap {
			seeds = extractSeeds(line)
		}
		if strings.Contains(line, "map:") && !expectMap {
			mapName = strings.Split(line, " ")[0]
			expectMap = true
		} else if len(line) != 0 && expectMap {
			o := extractDigits(line)
			dest = append(dest, o[0])
			source = append(source, o[1])
			length = append(length, o[2])

		}
		if len(line) == 0 && expectMap {
			var m mapper
			m.dest = dest
			m.source = source
			m.length = length
			h[mapName] = m
			dest = []int64{}
			source = []int64{}
			length = []int64{}
			expectMap = false
		}
	}

	var end int64 = math.MaxInt64
	found := false
	for j := int64(0); j < end; j++ {
		humidity := inverseGet(h["humidity-to-location"], j)
		temp := inverseGet(h["temperature-to-humidity"], humidity)
		light := inverseGet(h["light-to-temperature"], temp)
		water := inverseGet(h["water-to-light"], light)
		fertilizer := inverseGet(h["fertilizer-to-water"], water)
		soil := inverseGet(h["soil-to-fertilizer"], fertilizer)
		seed := inverseGet(h["seed-to-soil"], soil)
		for i := 0; i < len(seeds); i += 2 {
			seedStart := seeds[i]
			length := seeds[i+1]
			if within(seedStart, length, seed) {
				return j
			}
		}
		if found {
			break
		}
	}
	return 0

}

func task1(input string) int64 {
	h := make(map[string]mapper)
	scanner := bufio.NewScanner(strings.NewReader(input))
	expectMap := false
	dest := []int64{}
	source := []int64{}
	length := []int64{}
	seeds := []int64{}
	mapName := ""

	for scanner.Scan() { // internally, it advances token based on sperator
		line := scanner.Text()
		if strings.Contains(line, "seeds:") && !expectMap {
			seeds = extractSeeds(line)
		}
		if strings.Contains(line, "map:") && !expectMap {
			mapName = strings.Split(line, " ")[0]
			expectMap = true
		} else if len(line) != 0 && expectMap {
			o := extractDigits(line)
			dest = append(dest, o[0])
			source = append(source, o[1])
			length = append(length, o[2])

		}
		if len(line) == 0 && expectMap {
			var m mapper
			m.dest = dest
			m.source = source
			m.length = length
			h[mapName] = m
			dest = []int64{}
			source = []int64{}
			length = []int64{}
			expectMap = false
		}
	}

	location := []int64{}
	var lowest int64 = math.MaxInt64
	for _, seed := range seeds {
		soil := get(h["seed-to-soil"], seed)
		fertilizer := get(h["soil-to-fertilizer"], soil)
		water := get(h["fertilizer-to-water"], fertilizer)
		light := get(h["water-to-light"], water)
		temp := get(h["light-to-temperature"], light)
		humidity := get(h["temperature-to-humidity"], temp)
		loc := get(h["humidity-to-location"], humidity)
		location = append(location, loc)
		if loc < lowest {
			lowest = loc
		}
	}

	return lowest
}

func extractSeeds(line string) []int64 {
	digitStrings := strings.Split(line, ": ")[1]
	digits := strings.Split(digitStrings, " ")
	seeds := []int64{}
	for _, d := range digits {
		v, _ := strconv.ParseInt(d, 10, 64)
		seeds = append(seeds, v)
	}
	return seeds
}

func extractDigits(line string) []int64 {
	digits := strings.Split(line, " ")
	seeds := []int64{}
	for _, d := range digits {
		v, _ := strconv.ParseInt(d, 10, 64)
		seeds = append(seeds, v)
	}

	return seeds
}

func extractMap(line string) []int64 {
	digitStrings := strings.Split(line, ":")[1]
	digits := strings.Split(digitStrings, " ")
	seeds := []int64{}
	for _, d := range digits {
		v, _ := strconv.ParseInt(d, 10, 64)
		seeds = append(seeds, v)
	}
	return seeds
}

func get(m mapper, v int64) int64 {
	source := m.source
	dest := m.dest
	l := m.length
	for i, a := range source {
		if v >= a && v < a+l[i] {
			return dest[i] + (v - a)
		}
	}
	return v
}

func inverseGet(m mapper, v int64) int64 {
	source := m.source
	dest := m.dest
	l := m.length
	for i, a := range dest {
		if v >= a && v < a+l[i] {
			return source[i] + (v - a)
		}
	}
	return v
}

func within(s int64, l int64, v int64) bool {
	if v >= s && v < s+l {
		return true
	}
	return false
}
