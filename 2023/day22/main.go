package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Brick struct {
	name, x, y, z, w, h, d int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	bricks := Parse(input)
	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i].z < bricks[j].z
	})
	Settle(bricks)

	result := 0
	for _, v := range bricks {
		removed := RemoveBrick(*v, bricks)
		if HasSettled(removed) {
			result++
		}
	}
	return result
}

func RemoveBrick(brick Brick, bricks []*Brick) []*Brick {
	out := []*Brick{}
	for _, b := range bricks {
		if brick.name != b.name {
			out = append(out, b)
		}
	}
	return out
}

func InitMap(mx, my, mz int) [][][]int {
	out := [][][]int{}
	for z := 0; z <= mz; z++ {
		layer := [][]int{}
		for y := 0; y <= my; y++ {
			row := []int{}
			for x := 0; x <= mx; x++ {
				row = append(row, 0)
			}
			layer = append(layer, row)
		}
		out = append(out, layer)
	}
	return out
}

func Map(bricks []*Brick, Map [][][]int) [][][]int {
	for _, b := range bricks {
		for z := b.z; z < b.z+b.d; z++ {
			for y := b.y; y < b.y+b.h; y++ {
				for x := b.x; x < b.x+b.w; x++ {
					Map[z][y][x] = b.name
				}
			}
		}
	}
	return Map
}

func Settle(bricks []*Brick) {
	for !HasSettled(bricks) {
		for _, b := range bricks {
			if b.CanMove(bricks) {
				b.Move()
			}
		}
	}
}

func HasSettled(bricks []*Brick) bool {
	for _, b := range bricks {
		if b.CanMove(bricks) {
			return false
		}
	}
	return true
}

func task2(input string) int {
	return 0
}

func Parse(input string) []*Brick {
	out := []*Brick{}
	line := strings.Split(input, "\n")
	idx := 1
	for _, v := range line {
		p := strings.Split(v, "~")
		p1 := util.ToInt(strings.Split(p[0], ","))
		p2 := util.ToInt(strings.Split(p[1], ","))
		out = append(out, New(idx, p1[0], p1[1], p1[2], p2[0], p2[1], p2[2]))
		idx++
	}
	return out
}

func New(i, x1, y1, z1, x2, y2, z2 int) *Brick {
	return &Brick{
		name: i,
		x:    x1,
		y:    y1,
		z:    z1 - 1,
		w:    x2 + 1 - x1,
		h:    y2 + 1 - y1,
		d:    z2 + 1 - z1,
	}
}

func (brick *Brick) Move() {
	brick.z = brick.z - 1
}

func (brick Brick) CanMove(bricks []*Brick) bool {
	for _, b := range bricks {
		//skip the current brick
		if brick.name != b.name {
			// get the next z axis of the brick
			if (brick.z - 1) < 0 { // at bottom can't move
				return false
			}
			if (brick.z-1) <= (b.z+b.d-1) && !(b.z > brick.z-1) { // same z
				r1x1 := brick.x
				r1x2 := brick.x + brick.w
				r2x1 := b.x
				r2x2 := b.x + b.w
				r1y1 := brick.y
				r1y2 := brick.y + brick.h
				r2y1 := b.y
				r2y2 := b.y + b.h
				if DoesOverlap(r1x1, r1x2, r2x1, r2x2, r1y1, r1y2, r2y1, r2y2) {
					return false
				}
			}
		}
	}
	return true
}

func DoesOverlap(r1x1, r1x2, r2x1, r2x2, r1y1, r1y2, r2y1, r2y2 int) bool {
	if r1x1 > r2x2 || r1x2 < r2x1 ||
		r1y1 > r2y2 || r1y2 < r2y1 {
		return false
	}
	return true
}

func FindBoundries(bricks []*Brick) (int, int, int) {
	maxX := 0
	maxY := 0
	maxZ := 0
	for _, v := range bricks {
		if v.x+v.w > maxX {
			maxX = v.x + v.w - 1
		}
		if v.y+v.h > maxY {
			maxY = v.y + v.h - 1
		}
		if v.z+v.d > maxZ {
			maxZ = v.z + v.d - 1
		}
	}
	return maxX, maxY, maxZ
}

func PrintMapZX(m [][][]int, slice int) {
	fmt.Println("----XZ-----")
	for z := len(m) - 1; z >= 0; z-- {
		row := []rune{}
		for x := 0; x < len(m[0][0]); x++ {
			if m[z][slice][x] == 0 {
				row = append(row, '.')
			} else {
				row = append(row, 'X')
			}
		}
		fmt.Println(string(row))
	}
}

func PrintMapZY(m [][][]int, slice int) {
	fmt.Println("--YZ---")
	for z := len(m) - 1; z >= 0; z-- {
		row := []rune{}
		for y := 0; y < len(m[0]); y++ {
			if m[z][y][slice] == 0 {
				row = append(row, '.')
			} else {
				row = append(row, 'X')
			}
		}
		fmt.Println(string(row))
	}
}
