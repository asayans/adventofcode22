package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x    int
	y    int
	info *CoordInfo
}

type CoordInfo struct {
	height                int
	neighbourMinMaxHeight int
}

func read(file string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner, f
}

func (coord *Coordinate) printCoordInfo() {
	fmt.Printf("\n")
	fmt.Printf("%d, %d: \n\theight: %d \n\tneighbourMinMaxHeight: %d",
		coord.x, coord.y, coord.info.height, coord.info.neighbourMinMaxHeight)
	fmt.Printf("\n")
}

func maxIntSlice(s []int) int {
	m := 0
	for _, e := range s {
		if e > m {
			m = e
		}
	}
	return m
}

func traverseMatrix(matrix [][]int) [][]int {
	var traversed [][]int
	for j := 0; j < len(matrix); j++ {
		tmpSlice := []int{}
		for i := 0; i < len(matrix[0]); i++ {
			tmpSlice = append(tmpSlice, matrix[i][j])
		}
		traversed = append(traversed, tmpSlice)
	}
	return traversed
}

func findCoord(x int, y int, coords []*Coordinate) *Coordinate {
	for _, c := range coords {
		if c.x == x && c.y == y {
			return c
		}
	}
	return &Coordinate{}
}

func isHorizontalEdge(i int, s [][]int) bool {
	if i == 0 || i == len(s)-1 {
		return true
	}
	return false
}

func isVerticalEdge(i int, s []int) bool {
	if i == 0 || i == len(s)-1 {
		return true
	}
	return false
}

func processMatrix(matrix [][]int) []*Coordinate {
	var data []*Coordinate
	for i, line := range matrix {
		for j, h := range line {
			if isHorizontalEdge(i, matrix) {
				coord := Coordinate{x: i, y: j, info: &CoordInfo{height: h, neighbourMinMaxHeight: -1}}
				data = append(data, &coord)
				continue
			}
			if isVerticalEdge(j, line) {
				coord := Coordinate{x: i, y: j, info: &CoordInfo{height: h, neighbourMinMaxHeight: -1}}
				data = append(data, &coord)
				continue
			}
			coord := Coordinate{x: i, y: j, info: &CoordInfo{height: h, neighbourMinMaxHeight: 100}}
			leftView := matrix[i][:j]
			rightView := matrix[i][j+1:]
			// fmt.Println("(", i, j, ")", "line: ", line, " leftView: ", leftView, " rightView: ", rightView, "neighbourMinMaxHeight: ", coord.info.neighbourMinMaxHeight)
			if maxIntSlice(leftView) < coord.info.neighbourMinMaxHeight {
				coord.info.neighbourMinMaxHeight = maxIntSlice(leftView)
			}
			if maxIntSlice(rightView) < coord.info.neighbourMinMaxHeight {
				coord.info.neighbourMinMaxHeight = maxIntSlice(rightView)
			}
			coord.info.height = h
			data = append(data, &coord)
		}
	}
	traversedMatrix := traverseMatrix(matrix)
	for i, line := range traversedMatrix {
		if isHorizontalEdge(i, traversedMatrix) {
			continue
		}
		for j := range line {
			if isVerticalEdge(j, line) {
				continue
			}
			coord := findCoord(j, i, data)
			leftView := line[:j]
			rightView := line[j+1:]
			// fmt.Println("Traversed (", i, j, ")", "line: ", line, " leftView: ", leftView, " rightView: ", rightView)
			//coord.printCoordInfo()
			if maxIntSlice(leftView) < coord.info.neighbourMinMaxHeight {
				coord.info.neighbourMinMaxHeight = maxIntSlice(leftView)
			}
			if maxIntSlice(rightView) < coord.info.neighbourMinMaxHeight {
				coord.info.neighbourMinMaxHeight = maxIntSlice(rightView)
			}
		}
	}
	return data
}

func isVisible(tree *CoordInfo) bool {
	res := false
	if tree.height > tree.neighbourMinMaxHeight {
		res = true
	}
	return res
}

func main() {
	scanner, f := read("input.txt")
	defer f.Close()
	var matrix [][]int

	// Store input in matrix
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		treeSlice := []int{}
		for _, treeTxt := range line {
			tree, _ := strconv.Atoi(treeTxt)
			treeSlice = append(treeSlice, tree)
		}
		matrix = append(matrix, treeSlice)
	}
	// Process matrix and get relevant data
	data := processMatrix(matrix)
	visible := 0
	for _, c := range data {
		if isVisible(c.info) {
			visible++
		}
	}
	// findCoord(2, 1, data).printCoordInfo()
	fmt.Println("Visible trees: ", visible)
}
