package main

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type node struct {
	Name      string
	LeftNode  string
	RightNode string
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2, _ := task2(input)
	fmt.Printf("task2:%d\n", task2)
	task3, _ := task2BruteForce(input)
	fmt.Printf("task2-yolo:%d\n", task3)
}

func task1(input string) int {
	Directions, Nodes := ParseInput(input)
	found := false
	dirIdx := 0
	steps := 1
	CurrentNode := FindNode(Nodes, "AAA")
	for !found {
		if Directions[dirIdx] == "L" {
			CurrentNode = FindNode(Nodes, CurrentNode.LeftNode)
		} else {
			CurrentNode = FindNode(Nodes, CurrentNode.RightNode)
		}
		if CurrentNode.Name == "ZZZ" {
			found = true
			break
		}
		if dirIdx == len(Directions)-1 {
			dirIdx = 0
		} else {
			dirIdx++
		}
		steps++
	}
	return steps
}

func task2(input string) (int, error) {
	Directions, Nodes := ParseInput(input)
	final := 1
	results := []int{}
	startNodes := FindAllNodes(Nodes, "A")
	found := false
	dirIdx := 0
	steps := 1
	for _, n := range startNodes {
		startNode := n
		for !found {
			if Directions[dirIdx] == "L" {
				startNode = FindNode(Nodes, startNode.LeftNode)
			} else {
				startNode = FindNode(Nodes, startNode.RightNode)
			}
			if strings.HasSuffix(startNode.Name, "Z") {
				results = append(results, steps)
				steps = 0
				dirIdx = 0
				found = true
			}
			if dirIdx == len(Directions)-1 {
				dirIdx = 0
			} else {
				dirIdx++
			}
			steps++
		}
		found = false
		dirIdx = 0

	}
	gcds := []int{}
	for i, x := range results {
		for j, y := range results {
			if i != j {
				gcds = append(gcds, FindGcd(x, y))
			}
		}
	}
	v := gcds[0]
	gcdFound := true
	for i := 1; i < len(gcds); i++ {
		if v != gcds[i] {
			gcdFound = false
		}
	}
	if !gcdFound {
		return 0, errors.New("no common denominator found")
	}
	for i := 0; i < len(results); i++ {
		final = final * results[i] / gcds[0]
	}
	return final * gcds[0], nil
}

func task2BruteForce(input string) (int, error) {
	Directions, Nodes := ParseInput(input)
	startNodes := FindAllNodes(Nodes, "A")
	found := false
	dirIdx := 0
	steps := 1
	for !found {
		for i, n := range startNodes {
			if Directions[dirIdx] == "L" {
				startNodes[i] = FindNode(Nodes, n.LeftNode)
			} else {
				startNodes[i] = FindNode(Nodes, n.RightNode)
			}
		}
		if AllEndsWith(startNodes, "Z") {
			found = true
			break
		}
		if dirIdx == len(Directions)-1 {
			dirIdx = 0
		} else {
			dirIdx++
		}
		steps++
	}
	return steps, nil
}

func ParseInput(input string) (DirectionList []string, nodes []node) {
	TempNodes := []node{}
	TempDirectionList := []string{}
	re := regexp.MustCompile(`(?m)(\w+) = \((\w+), (\w+)\)`)
	scanner := bufio.NewScanner(strings.NewReader(input))
	line := 0
	for scanner.Scan() {
		if line == 0 {
			TempDirectionList = append(TempDirectionList, strings.Split(scanner.Text(), "")...)
		}
		if line > 1 {
			val := re.FindAllStringSubmatch(scanner.Text(), -1)
			n := node{
				Name:      val[0][1],
				LeftNode:  val[0][2],
				RightNode: val[0][3],
			}
			TempNodes = append(TempNodes, n)
		}
		line++
	}
	return TempDirectionList, TempNodes
}

func FindNode(c []node, name string) (n node) {
	for _, n := range c {
		if n.Name == name {
			return n
		}
	}
	return node{}
}

func FindAllNodes(c []node, nameEndsWith string) (n []node) {
	tempNodes := []node{}
	for _, n := range c {
		if strings.HasSuffix(n.Name, nameEndsWith) {
			tempNodes = append(tempNodes, n)
		}
	}
	return tempNodes
}

func AllEndsWith(c []node, nameEndsWith string) bool {
	found := 0
	for _, n := range c {
		if strings.HasSuffix(n.Name, nameEndsWith) {
			found++
		}
	}
	return found == len(c)
}

func FindGcd(a int, b int) int {
	for b != 0 {
		r := a % b
		a = b
		b = r
	}
	return a
}
