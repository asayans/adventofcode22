package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coordinate struct {
	x    int
	y    int
	info *CoordInfo
}

type CoordInfo struct {
	height      int
	top         *Coordinate
	bottom      *Coordinate
	left        *Coordinate
	right       *Coordinate
	scenicScore int
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
	fmt.Printf("%d, %d: \n\theight: %d \n\ttop: %v \n\tbottom: %v \n\tleft: %v \n\tright: %v \n\tscenicScore: %d",
		coord.x, coord.y, coord.info.height, coord.info.top, coord.info.bottom, coord.info.left, coord.info.right, coord.info.scenicScore)
	fmt.Printf("\n")
}

func findCoord(x int, y int, data [][]*Coordinate) *Coordinate {
	for _, line := range data {
		for _, c := range line {
			if c.x == x && c.y == y {
				return c
			}
		}
	}
	return &Coordinate{}
}

func updateCoordInfo(c *Coordinate, data [][]*Coordinate) {
	edgeLimit := len(data) - 1
	// By default, all trees can see to the edges
	c.info.top = findCoord(0, c.y, data)
	c.info.bottom = findCoord(edgeLimit, c.y, data)
	c.info.left = findCoord(c.x, 0, data)
	c.info.right = findCoord(c.x, edgeLimit, data)
	// Find tree with same height or taller in each side
	for x := c.x - 1; x >= 0; x-- {
		p := findCoord(x, c.y, data)
		if p.info.height >= c.info.height {
			c.info.top = p
			break
		}
	}
	for x := c.x + 1; x <= edgeLimit; x++ {
		p := findCoord(x, c.y, data)
		if p.info.height >= c.info.height {
			c.info.bottom = p
			break
		}
	}
	for y := c.y - 1; y >= 0; y-- {
		p := findCoord(c.x, y, data)
		if p.info.height >= c.info.height {
			c.info.left = p
			break
		}
	}
	for y := c.y + 1; y <= edgeLimit; y++ {
		p := findCoord(c.x, y, data)
		if p.info.height >= c.info.height {
			c.info.right = p
			break
		}
	}
}

func addCoordsInfo(data [][]*Coordinate) {
	edgeLimit := len(data) - 1
	for _, line := range data {
		for _, c := range line {
			if c.x != 0 && c.y != 0 && c.y != edgeLimit && c.x != edgeLimit {
				updateCoordInfo(c, data)
			} else {
				// Edge tree, one side will be 0, so scenic score = 0. Using -1 as control
				c.info.scenicScore = -1
			}
		}
	}
}

func dist(t int, p int) int {
	return int(math.Abs(float64(t) - float64(p)))
}

func scenicScore(coord *Coordinate) int {
	scenicScore := dist(coord.y, coord.info.left.y) * dist(coord.y, coord.info.right.y) *
		dist(coord.x, coord.info.top.x) * dist(coord.x, coord.info.bottom.x)
	return scenicScore
}

func main() {
	start := time.Now()
	scanner, f := read("input.txt")
	defer f.Close()
	i := 0
	data := [][]*Coordinate{}
	// Store input in data
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		treeSlice := []*Coordinate{}
		for j, treeTxt := range line {
			h, _ := strconv.Atoi(treeTxt)
			coord := Coordinate{x: i, y: j, info: &CoordInfo{height: h}}
			treeSlice = append(treeSlice, &coord)
		}
		data = append(data, treeSlice)
		i++
	}
	// Add CoordInfo to every Coordinate in the forest
	addCoordsInfo(data)

	currentBest := 0
	var bestTree *Coordinate
	for _, line := range data {
		for _, c := range line {
			if c.info.scenicScore != -1 && scenicScore(c) > currentBest {
				currentBest = scenicScore(c)
				bestTree = c
			}
		}
	}
	// findCoord(2, 1, data).printCoordInfo()
	bestTree.printCoordInfo()
	fmt.Println("Highest scenic score: ", currentBest)

	// Execution time
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
