package main

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type Point struct {
	x int
	y int
	z int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input, int64(200000000000000), int64(400000000000000))
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input, math.MinInt64, math.MaxInt64)
	fmt.Printf("task2:%d\n", task2)
}

func CheckIntecept(a, b []Point) bool {
	m1, c1 := FindLine(a[0], a[1])
	m2, c2 := FindLine(b[0], b[1])
	_, _, e := FindIntecept(m1, c1, m2, c2)
	if e == nil {
		return true
	}
	return false
}

func task2(input string, start, end int64) int {
	v := Parse(input)
	maxInt, x, y, z := 0, 0, 0, 0
	for i := -250; i <= 250; i++ {
		for j := -250; j <= 250; j++ {
			v1 := AdjustVelocity(i, j, 0, v)
			intecepts := TotalIntecepts(v1, start, end)
			if intecepts > maxInt {
				maxInt = intecepts
				x = i
				y = j
			}
		}
	}
	switched := SwitchYZ(v)
	maxInt = 0
	for i := -250; i <= 250; i++ {
		for j := -250; j <= 250; j++ {
			v1 := AdjustVelocity(i, j, 0, switched)
			intecepts := TotalIntecepts(v1, start, end)
			if intecepts > maxInt {
				maxInt = intecepts
				x = i
				z = j
			}
		}
	}
	px, py, pz := float64(0), float64(0), float64(0)
	//	fmt.Printf("%d,%d,%d\n", maxInt, x, z)
	adjustedXY := AdjustVelocity(x, y, z, v)
	m1, c1 := FindLine(adjustedXY[0][0], adjustedXY[0][1])
	m2, c2 := FindLine(adjustedXY[1][0], adjustedXY[1][1])
	px, py, _ = FindIntecept(m1, c1, m2, c2)
	//	fmt.Printf("%f,%f\n", px, py)

	adjustedYZ := AdjustVelocity(x, z, y, switched)
	m1, c1 = FindLine(adjustedYZ[0][0], adjustedYZ[0][1])
	m2, c2 = FindLine(adjustedYZ[1][0], adjustedYZ[1][1])
	px, pz, _ = FindIntecept(m1, c1, m2, c2)
	//fmt.Printf("%f,%f\n", px, pz)

	return int(px + py + pz)
}

func task1(input string, start, end int64) int {
	v := Parse(input)
	return TotalIntecepts(v, start, end)
}

func AdjustVelocity(x, y, z int, v [][]Point) (out [][]Point) {
	out = [][]Point{}
	for i := 0; i < len(v); i++ {
		p2 := v[i][1]
		p := Point{x: p2.x + x, y: p2.y + y, z: p2.z + z}
		out = append(out, []Point{v[i][0], p})
	}
	return out
}

func SwitchYZ(v [][]Point) (out [][]Point) {
	out = [][]Point{}
	for i := 0; i < len(v); i++ {
		p1 := v[i][0]
		p2 := v[i][1]
		out = append(out, []Point{{x: p1.x, y: p1.z, z: p1.y}, {x: p2.x, y: p2.z, z: p2.y}})
	}
	return out
}

func SwitchXZ(v [][]Point) (out [][]Point) {
	out = [][]Point{}
	for i := 0; i < len(v); i++ {
		p1 := v[i][0]
		p2 := v[i][1]
		out = append(out, []Point{{x: p1.z, y: p1.y, z: p1.x}, {x: p2.z, y: p2.y, z: p2.x}})
	}
	return out
}

func FindVelocitiyBoundries(v [][]Point) (maxX, minX, maxY, minY, maxZ, minZ int) {
	minX, minY, minX = math.MaxInt64, math.MaxInt64, math.MaxInt64
	maxX, maxY, maxX = math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(v); i++ {
		x := v[i][1].x - v[i][0].x
		y := v[i][1].y - v[i][0].y
		z := v[i][1].z - v[i][0].z
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
		if z > maxZ {
			maxZ = z
		}
		if z < minZ {
			minZ = z
		}
	}
	return maxX, minX, maxY, minY, maxZ, minZ
}

func TotalIntecepts(v [][]Point, start, end int64) int {
	result := 0
	for i := 0; i < len(v); i++ {
		m1, c1 := FindLine(v[i][0], v[i][1])
		for j := i; j < len(v); j++ {
			m2, c2 := FindLine(v[j][0], v[j][1])
			if i != j {
				x, y, e := FindIntecept(m1, c1, m2, c2)
				if e != nil {
					//fmt.Println("no intecept")
				} else {
					if x <= float64(end) && x >= float64(start) && y <= float64(end) && y >= float64(start) {
						p1xd := v[i][1].x - v[i][0].x
						p2xd := v[j][1].x - v[j][0].x
						p1yd := v[i][1].y - v[i][0].y
						p2yd := v[j][1].y - v[j][0].y

						if p1xd < 0 && x > float64(v[i][0].x) {
							continue
						}
						if p2xd < 0 && x > float64(v[j][0].x) {
							continue
						}
						if p1yd < 0 && y > float64(v[i][1].y) {
							continue
						}
						if p2yd < 0 && y > float64(v[j][1].y) {
							continue
						}
						if p1xd > 0 && x < float64(v[i][0].x) {
							continue
						}
						if p2xd > 0 && x < float64(v[j][0].x) {
							continue
						}
						if p1yd > 0 && y < float64(v[i][1].y) {
							continue
						}
						if p2yd > 0 && y < float64(v[j][1].y) {
							continue
						}

						result++
					}
				}
			}
		}
	}
	return result
}

func Parse(input string) [][]Point {
	out := [][]Point{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		s := strings.Split(strings.ReplaceAll(line, " ", ""), "@")
		leftP := util.ToInt(strings.Split(s[0], ","))
		rightV := util.ToInt(strings.Split(s[1], ","))
		out = append(out, []Point{
			{leftP[0], leftP[1], leftP[2]},
			{leftP[0] + rightV[0], leftP[1] + rightV[1], leftP[2] + rightV[2]}})
	}
	return out
}

func FindLine(p1, p2 Point) (m, c float64) {
	m = (float64(p2.y) - float64(p1.y)) / (float64(p2.x) - float64(p1.x))
	c = ((float64(p2.y) + float64(p1.y)) - (m * (float64(p2.x) + float64(p1.x)))) / 2
	return m, c
}

func FindIntecept(m1, c1, m2, c2 float64) (x, y float64, e error) {
	if m1 == m2 {
		return 0, 0, errors.New("doesn't intercept")
	}
	x = (c2 - c1) / (m1 - m2)
	y = (m1 * (c2 - c1) / (m1 - m2)) + c1
	return x, y, nil
}
