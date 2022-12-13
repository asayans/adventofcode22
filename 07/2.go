package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func read() (*bufio.Scanner, *os.File) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner, f
}

type FileSystem struct {
	root *Dir
}

func newFileSystem() *FileSystem {
	return &FileSystem{
		root: nil,
	}
}

type File struct {
	name string
	size int
}

func newFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

type Dir struct {
	name        string
	parent      *Dir
	directories []*Dir
	files       []*File
	size        int
}

func newDir(name string) *Dir {
	return &Dir{
		name:        name,
		parent:      nil,
		directories: nil,
		files:       nil,
		size:        0,
	}
}

func (parent *Dir) addDir(child *Dir) {
	parent.directories = append(parent.directories, child)
}

func (dir *Dir) addFile(file *File) {
	dir.files = append(dir.files, file)
}

func (dir *Dir) getDirNames() []string {
	var l []string
	for _, d := range dir.directories {
		l = append(l, d.name)
	}
	return l
}

func (dir *Dir) getFileNames() []string {
	var l []string
	for _, d := range dir.files {
		l = append(l, d.name)
	}
	return l
}

func (dir *Dir) getFileSizes() []int {
	var l []int
	for _, d := range dir.files {
		l = append(l, d.size)
	}
	return l
}

func (dir *Dir) findDirByName(name string) *Dir {
	for _, d := range dir.directories {
		if d.name == name {
			return d
		}
	}
	return newDir("")
}

func (dir *Dir) printDir() {
	fmt.Printf("\n")
	fmt.Printf("%s: \n\tparent: %s \n\tdirectories: %v \n\tfiles: %v\n\tfileSizes: %v\n\tdirSize: %d",
		dir.name, dir.parent.name, dir.getDirNames(), dir.getFileNames(), dir.getFileSizes(), dir.size)
	fmt.Printf("\n")
}

func calculSize(dir *Dir) int {
	size := 0
	for _, f := range dir.files {
		size += f.size
	}
	for _, d := range dir.directories {
		size += calculSize(d)
	}
	dir.size = size
	return size
}

func totalSize(root *Dir, c chan *Dir) {
	for _, d := range root.directories {
		c <- d
		totalSize(d, c)
	}

}

func main() {
	scanner, f := read()
	defer f.Close()

	currentDir := newDir("/")
	currentDir.parent = newDir("")
	fs := newFileSystem()
	fs.root = currentDir
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if line[0] == "$" {
			// Directories tree creation
			if line[1] == "cd" && line[2] == ".." {
				currentDir = currentDir.parent
			} else if line[1] == "cd" && line[2] != "/" {
				tmpDir := newDir(line[2])
				currentDir.addDir(tmpDir)
				tmpDir.parent = currentDir
				currentDir = tmpDir
			}
		} else { // We are reading the output of 'ls' command
			if line[0] != "dir" {
				size, _ := strconv.Atoi(line[0])
				file := newFile(line[1], size)
				currentDir.addFile(file)
			}
		}
	}

	// File scanned
	calculSize(fs.root)
	c := make(chan *Dir)

	go func() {
		defer close(c)
		totalSize(fs.root, c)
	}()

	totalSpace := 70000000
	freeSpace := totalSpace - fs.root.size
	updateSize := 30000000
	neededSpace := updateSize - freeSpace
	currentDirSize := totalSpace
	var res int
	for d := range c {
		if d.size >= neededSpace && d.size <= currentDirSize {
			res = d.size
			currentDirSize = d.size
		}
	}
	fmt.Println(res)
}
