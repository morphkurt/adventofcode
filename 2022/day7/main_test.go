package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var i = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestTask1(t *testing.T) {
	got := task1(i)
	want := 95437

	assert.Equal(t, want, got)
}

func TestTask2(t *testing.T) {
	got := task2(i)
	want := 24933642

	assert.Equal(t, want, got)
}
