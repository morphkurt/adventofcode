package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(f string) string {
	b, err := os.ReadFile(f) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	return str
}

func Transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func ToInt(slice []string) []int {
	out := []int{}
	for _, v := range slice {
		s := strings.TrimSpace(v)
		i, _ := strconv.Atoi(s)
		out = append(out, i)
	}
	return out
}

func ToUint(slice []string) []uint {
	out := []uint{}
	for _, v := range slice {
		s := strings.TrimSpace(v)
		i, _ := strconv.Atoi(s)
		out = append(out, uint(i))
	}
	return out
}
