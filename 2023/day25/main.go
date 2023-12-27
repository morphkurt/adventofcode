package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Node struct {
	name      string
	connected []*Node
	distance  map[string]int
}

type GraphNode struct {
	node    *Node
	visited []*Node
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input, 100)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func (node *Node) RemoveNode(n string) {
	out := []*Node{}
	for _, v := range node.connected {
		if v.name != n {
			out = append(out, v)
		}
	}
	node.connected = out
}

func task1(input string, sample int) int {
	list, m := FreshNodes(input)
	nodes := [][]*Node{}
	for j := 1; j < sample; j++ {
		n := Map(list[j], list[len(list)-j-1])
		nodes = append(nodes, n...)
	}

	_, keys := CountPairs(nodes)
	//find top 3 pairs

	pair1 := strings.Split(keys[0], ",")
	pair2 := strings.Split(keys[1], ",")
	pair3 := strings.Split(keys[2], ",")

	fmt.Println(pair1)
	fmt.Println(pair2)
	fmt.Println(pair3)

	m[pair1[0]].RemoveNode(pair1[1])
	m[pair1[1]].RemoveNode(pair1[0])
	m[pair2[0]].RemoveNode(pair2[1])
	m[pair2[1]].RemoveNode(pair2[0])
	m[pair3[0]].RemoveNode(pair3[1])
	m[pair3[1]].RemoveNode(pair3[0])
	groups := 0
	for j := 1; j < len(list); j++ {
		n := Map(list[0], list[j])
		if len(n) > 0 {
			groups++
		}
	}

	return (groups * (len(m) - groups))
}

func FreshNodes(input string) ([]*Node, map[string]*Node) {
	m := Parse(input)
	l := []*Node{}
	//	fmt.Println(v)
	for _, v := range m {
		l = append(l, v)
	}
	return l, m
}

func task2(input string) int {
	return 0
}

func remove(slice []*Node, s int) []*Node {
	return append(slice[:s], slice[s+1:]...)
}

func CountPairs(n [][]*Node) (map[string]int, []string) {
	out := map[string]int{}
	for _, v := range n {
		for i := 0; i < len(v)-1; i = i + 1 {
			comp := []string{v[i].name, v[i+1].name}
			sort.Strings(comp)
			key := fmt.Sprintf("%s,%s", comp[0], comp[1])
			if val, ok := out[key]; ok {
				out[key] = val + 1
			} else {
				out[key] = 1
			}
		}
	}
	keys := make([]string, 0, len(out))
	for key := range out {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return out[keys[i]] > out[keys[j]] })

	return out, keys
}

func Map(curr, end *Node) [][]*Node {
	out := [][]*Node{}
	queue := []GraphNode{}
	visitedNodes := []*Node{}
	for _, v := range curr.connected {
		queue = append(queue, GraphNode{node: v, visited: []*Node{}})
	}
	for len(queue) > 0 {
		n := queue[0]
		visitedNodes = append(visitedNodes, n.node)
		queue = queue[1:]
		v := []*Node{}
		v = append(v, n.visited...)
		v = append(v, n.node)
		if Contains(n.node, n.visited) {
			continue
		}
		if end.name == n.node.name {
			out = append(out, v)
			return out
		} else {
			for _, e := range n.node.connected {
				if !Contains(e, visitedNodes) {
					queue = append(queue, GraphNode{node: e, visited: v})
				}
			}
		}
	}
	return out
}

func Contains(node *Node, nodes []*Node) bool {
	for _, v := range nodes {
		if v.name == node.name {
			return true
		}
	}
	return false
}

func Parse(input string) map[string]*Node {
	out := map[string]*Node{}
	for _, line := range strings.Split(input, "\n") {
		left := strings.Split(line, ": ")[0]
		leftNode, ok := out[left]
		if !ok {
			leftNode = &Node{name: left, connected: []*Node{}, distance: map[string]int{}}
		}
		right := strings.Split(line, ": ")[1]
		connected := strings.Split(right, " ")
		for _, n := range connected {
			node, ok := out[n]
			if !ok {
				node = &Node{name: n, connected: []*Node{}, distance: map[string]int{}}
			}
			node.connected = append(node.connected, leftNode)
			leftNode.connected = append(leftNode.connected, node)
			out[n] = node
		}
		out[left] = leftNode
	}
	return out
}
