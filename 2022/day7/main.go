package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/morphkurt/adventofcode/util"
)

type directory struct {
	name        string
	parent      *directory
	directories []*directory
	files       []*file
}

const (
	ROOT      = "/"
	MOVE_BACK = ".."
)

type file struct {
	name string
	size int
}

func main() {
	input := util.ReadFile("input")
	task1 := task1(input)
	fmt.Printf("task1:%d\n", task1)
	task2 := task2(input)
	fmt.Printf("task2:%d\n", task2)
}

func task1(input string) int {
	root := parse(input)
	result := 0
	queue := []*directory{}
	queue = append(queue, root)

	for len(queue) > 0 {
		dir := queue[0]
		queue = queue[1:]
		size := dir.getSize()
		if size < 100000 {
			result += size
		}
		queue = append(queue, dir.directories...)
	}

	return result
}

func task2(input string) int {
	root := parse(input)
	usedSpace := root.getSize()
	freeSpace := 70000000 - usedSpace
	spaceNeededForUpdate := 30000000 - freeSpace
	toDelete := []int{}
	queue := []*directory{}
	queue = append(queue, root)

	for len(queue) > 0 {
		dir := queue[0]
		queue = queue[1:]
		size := dir.getSize()
		if size > spaceNeededForUpdate {
			toDelete = append(toDelete, size)
		}
		queue = append(queue, dir.directories...)
	}
	sort.Ints(toDelete)
	return toDelete[0]
}

func (dir directory) getSize() int {
	fileSize := 0
	for _, v := range dir.files {
		fileSize += v.size
	}
	childDirSize := 0
	for _, v := range dir.directories {
		childDirSize += v.getSize()
	}
	return fileSize + childDirSize
}

func parse(input string) *directory {
	root := &directory{
		name:        ROOT,
		directories: []*directory{},
		files:       []*file{},
	}
	currentDir := root
	lines := strings.Split(input, "\n")
	expectListing := false
	for _, line := range lines {
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				expectListing = false
				var name string
				fmt.Sscanf(line, "$ cd %s", &name)
				switch name {
				case MOVE_BACK:
					currentDir = currentDir.parent
				case ROOT:
					currentDir = root
				default:
					currentDir = getDir(name, currentDir.directories)
				}
			} else {
				expectListing = true
			}
		}
		if expectListing && !strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "dir") {
				dirName := ""
				fmt.Sscanf(line, "dir %s", &dirName)
				dir := &directory{
					parent:      currentDir,
					name:        dirName,
					directories: []*directory{},
					files:       []*file{},
				}
				currentDir.directories = append(currentDir.directories, dir)
			} else {
				size := 0
				name := ""
				fmt.Sscanf(line, "%d %s", &size, &name)
				f := &file{
					name: name,
					size: size,
				}
				currentDir.files = append(currentDir.files, f)
			}
		}
	}
	return root
}

func getDir(name string, dirs []*directory) *directory {
	for _, v := range dirs {
		if v.name == name {
			return v
		}
	}
	return nil
}
