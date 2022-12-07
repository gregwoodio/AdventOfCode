package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	parent  *dir
	path    string
	name    string
	files   []*file
	subdirs []*dir
	size    int
}

func solvePartOne(input []string) int {
	root := mkfs(input)
	getDirectorySize(root)

	sum := 0
	largeDirs := []*dir{}
	largeDirs = *(getDirsComparedBy(root, &largeDirs, dirSizeLessThan, 100000))
	for _, d := range largeDirs {
		sum += d.size
	}

	return sum
}

func solvePartTwo(input []string) int {
	total := 70000000
	toFree := 30000000

	root := mkfs(input)
	getDirectorySize(root)

	used := total - root.size
	target := toFree - used

	largeDirs := []*dir{}
	largeDirs = *(getDirsComparedBy(root, &largeDirs, dirSizeGreaterThan, target))

	// find smallest of the dirs big enough to delete
	smallest := total
	for _, d := range largeDirs {
		if d.size < smallest {
			smallest = d.size
		}
	}

	return smallest
}

func getDirsComparedBy(pwd *dir, list *[]*dir, comp func(*dir, int) bool, n int) *[]*dir {
	if comp(pwd, n) {
		l := append(*list, pwd)
		list = &l
	}

	for _, subdir := range pwd.subdirs {
		list = getDirsComparedBy(subdir, list, comp, n)
	}
	return list
}

func dirSizeLessThan(d *dir, n int) bool {
	return d.size <= n
}

func dirSizeGreaterThan(d *dir, n int) bool {
	return d.size > n
}

func getDirectorySize(pwd *dir) {
	size := 0
	for _, f := range pwd.files {
		size += f.size
	}

	for _, d := range pwd.subdirs {
		if d.size == -1 {
			getDirectorySize(d)
		}
		size += d.size
	}

	pwd.size = size
}

// make file system, return root directory
func mkfs(input []string) *dir {
	dirs := map[string]*dir{}
	var pwd *dir

	for i := 0; i < len(input); i++ {
		if input[i][0] == '$' {
			// commands
			parts := strings.Split(input[i], " ")
			if parts[1] == "cd" {
				if parts[2] == ".." {
					if pwd.parent != nil {
						pwd = pwd.parent
					}
				} else {
					var path string
					if pwd != nil {
						path = fmt.Sprintf("%s/%s", pwd.path, parts[2])
					} else {
						path = "/"
					}
					if _, ok := dirs[path]; !ok {
						dirs[path] = mkdir(parts[2], pwd)
					}
					pwd = dirs[path]
				}
			} else if parts[1] == "ls" {
				i++
				for {
					if i >= len(input) {
						break
					}

					if input[i][0] == '$' {
						// new command, we're done reading ls output
						i--
						break
					}

					parts := strings.Split(input[i], " ")
					path := fmt.Sprintf("%s/%s", pwd.path, parts[1])
					if parts[0] == "dir" {
						if _, ok := dirs[path]; !ok {
							dirs[path] = mkdir(parts[1], pwd)
						}

						d := dirs[path]
						pwd.subdirs = append(pwd.subdirs, d)
					} else {
						// files
						size, err := strconv.Atoi(parts[0])
						if err != nil {
							fmt.Printf("Could not parse '%s' to int\n", parts[0])
						}
						pwd.files = append(pwd.files, &file{
							name: parts[1],
							size: size,
						})
					}
					i++
				}
			}
		}
	}

	return dirs["/"]
}

func mkdir(name string, parent *dir) *dir {
	var path string
	if parent == nil {
		path = "/"
	} else {
		path = fmt.Sprintf("%s/%s", parent.path, name)
	}

	return &dir{
		name:    name,
		path:    path,
		parent:  parent,
		files:   []*file{},
		subdirs: []*dir{},
		size:    -1,
	}
}
