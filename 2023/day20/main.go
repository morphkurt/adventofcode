package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

const (
	FLIPFLOP    = iota
	CONJUNCTION = iota
	BUTTON      = iota
	BROADCASTER = iota
	HIGH        = iota
	LOW         = iota
	VOID        = iota
)

type Signal struct {
	Type int
}

type BaseModule struct {
	Buffer    []map[string]Signal
	Name      string
	Outputs   []*Module
	HighCount int
	LowCount  int
}

type Broadcaster struct {
	BaseModule
}

type Flipflop struct {
	On bool
	BaseModule
}

type Conjuction struct {
	BaseModule
	Memory map[string]Signal
}

type Module interface {
	buff(string, Signal)
	process()
	set(*Module)
	getCount() (int, int)
	getName() string
	getPending() int
	getOutputs() []*Module
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task2(input string) int {
	m, o := Parse(input)
	moduleMap := Init(m, o)
	b := moduleMap["broadcaster"]
	count := 1000000
	kh, lz, tg, hn := false, false, false, false
	iter := []int{}
	for i := 1; i < count; i++ {
		b.buff("button", Signal{Type: LOW})
		pending := 1
		for pending != 0 {

			pending = 0
			for _, v := range moduleMap {
				v.process()
			}
			for _, v := range moduleMap {
				if v.getPending() > 0 {
					pending++
				}
			}
		}

		nextkh := moduleMap["kh"].(*Conjuction).HighCount
		nextlz := moduleMap["lz"].(*Conjuction).HighCount
		nexttg := moduleMap["tg"].(*Conjuction).HighCount
		nexthn := moduleMap["hn"].(*Conjuction).HighCount

		if nextkh == 1 && !kh {
			iter = append(iter, i)
			kh = true
		}
		if nextlz == 1 && !lz {
			iter = append(iter, i)
			lz = true
		}
		if nexttg == 1 && !tg {
			iter = append(iter, i)
			tg = true
		}
		if nexthn == 1 && !hn {
			iter = append(iter, i)
			hn = true
		}
		if len(iter) == 4 {
			break
		}
	}

	return iter[0] * iter[1] * iter[2] * iter[3]
}

func task1(input string) int {
	m, o := Parse(input)
	moduleMap := Init(m, o)
	b := moduleMap["broadcaster"]
	count := 1000
	for i := 1; i < count; i++ {
		b.buff("button", Signal{Type: LOW})
		pending := 1
		for pending != 0 {

			pending = 0
			for _, v := range moduleMap {
				v.process()
			}
			for _, v := range moduleMap {
				if v.getPending() > 0 {
					pending++
				}
			}
		}
	}
	highCount := 0
	lowCount := count

	for _, v := range moduleMap {
		h, l := v.getCount()
		highCount += h
		lowCount += l
	}
	return highCount * lowCount
}

func (f *Flipflop) process() {
	for _, b := range f.Buffer {
		for _, v := range b {
			t := v.Type
			r := Signal{Type: VOID}
			if f.On == false && t == HIGH {
				/// do nothing
			} else if f.On == true && t == LOW {
				r.Type = LOW
				f.On = false
			} else if f.On == false && t == LOW {
				r.Type = HIGH
				f.On = true
			}
			if r.Type != VOID {
				for _, v := range f.Outputs {
					m := *v
					if r.Type == HIGH {
						f.HighCount++
					} else {
						f.LowCount++
					}
					m.buff(f.Name, r)
				}
			}
		}
	}
	f.Buffer = []map[string]Signal{}
}

func (c *Conjuction) process() {
	for _, b := range c.Buffer {
		for k, t := range b {
			if t.Type != VOID {
				c.Memory[k] = t
			}
			r := Signal{Type: VOID}
			foundHigh := true
			for _, v := range c.Memory {
				if v.Type != HIGH {
					foundHigh = false
					break
				}
			}
			if foundHigh {
				r = Signal{Type: LOW}
			} else {
				r = Signal{Type: HIGH}
			}
			if t.Type != VOID {
				for _, v := range c.Outputs {
					m := *v
					if r.Type == HIGH {
						c.HighCount++
					} else {
						c.LowCount++
					}
					m.buff(c.Name, r)
				}
			}
		}
	}
	c.Buffer = []map[string]Signal{}
}

func (c *Broadcaster) process() {
	for _, b := range c.Buffer {
		for _, v := range b {
			for _, o := range c.Outputs {
				if v.Type == HIGH {
					c.HighCount++
				} else {
					c.LowCount++
				}
				m := *o
				m.buff(c.Name, v)
			}
		}
	}
	c.Buffer = []map[string]Signal{}
}

func (b *BaseModule) set(m *Module) {
	b.Outputs = append(b.Outputs, m)
}

func (b *BaseModule) buff(name string, s Signal) {
	b.Buffer = append(b.Buffer, map[string]Signal{
		name: s,
	})
}

func (b *BaseModule) getCount() (int, int) {
	h := b.HighCount
	l := b.LowCount
	return h, l
}

func (b *BaseModule) getPending() int {
	return len(b.Buffer)
}

func (b *BaseModule) getOutputs() []*Module {
	return b.Outputs
}

func (b *BaseModule) getName() string {
	return b.Name
}

func (b *Conjuction) getMemory() string {
	mem := b.Memory
	out := ""
	for _, v := range mem {
		out += strconv.Itoa(v.Type)
	}
	return out
}

func Init(names, outputs [][]string) map[string]Module {
	moduleMap := make(map[string]Module, 0)
	for _, v := range names {
		switch v[1] {
		case "F":
			moduleMap[v[0]] = &Flipflop{
				BaseModule: BaseModule{
					Name:    v[0],
					Outputs: []*Module{},
				},
				On: false,
			}
		case "C":
			memory := make(map[string]Signal)
			inputs := FindInputsForConjunction(v[0], names, outputs)
			for _, i := range inputs {
				memory[i] = Signal{Type: LOW}
			}
			moduleMap[v[0]] = &Conjuction{
				BaseModule: BaseModule{
					Name:    v[0],
					Outputs: []*Module{},
				},
				Memory: memory,
			}
		case "B":
			moduleMap[v[0]] = &Broadcaster{
				BaseModule: BaseModule{
					Name:    v[0],
					Outputs: []*Module{},
				},
			}
		}

	}
	for i := 0; i < len(names); i++ {
		m := moduleMap[names[i][0]]
		outputs := outputs[i]
		for _, v := range outputs {
			outputModule := moduleMap[v]
			m.set(&outputModule)
		}
		moduleMap[names[i][0]] = m
	}

	return moduleMap
}

func Parse(input string) ([][]string, [][]string) {
	modules := [][]string{}
	outputs := [][]string{}
	line := strings.Split(input, "\n")
	for _, l := range line {
		output := []string{}
		c := strings.Split(l, "->")
		moduleName := strings.Trim(c[0], " ")
		switch moduleName[0] {
		case '%':
			modules = append(modules, []string{moduleName[1:], "F"})
		case '&':
			modules = append(modules, []string{moduleName[1:], "C"})
		default:
			modules = append(modules, []string{moduleName, "B"})
		}
		outputArray := strings.Split(c[1], ",")
		for _, v := range outputArray {
			e := strings.Trim(v, " ")
			if e != "" {
				output = append(output, e)
			}
		}
		outputs = append(outputs, output)
	}
	return modules, outputs
}

func getIndex(e string, c [][]string) int {
	for i, v := range c[0] {
		if v == e {
			return i
		}
	}
	return -1
}

func FindInputsForConjunction(e string, names [][]string, outputs [][]string) []string {
	out := []string{}
	for i, v := range names {
		name := v[0]
		for _, m := range outputs[i] {
			if m == e {
				out = append(out, name)
			}
		}
	}
	return out
}
