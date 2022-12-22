package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Knot struct {
	x         int
	y         int
	direction string
}

func read(file string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner, f
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func (knot *Knot) move(dir string) {
	knot.direction = dir
	switch dir {
	case "U":
		knot.y = knot.y - 1
	case "D":
		knot.y = knot.y + 1
	case "R":
		knot.x = knot.x + 1
	case "L":
		knot.x = knot.x - 1
	}
}

func dist(p1 *Knot, p2 Knot) int {
	return int(math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2)))
}

func (knot *Knot) follow(followedKnot Knot) {
	opposite := make(map[string]string)
	opposite["U"] = "D"
	opposite["D"] = "U"
	opposite["L"] = "R"
	opposite["R"] = "L"
	// Does the knot need to follow?
	if knot.x == followedKnot.x && knot.y == followedKnot.y {
		log.Println("INFO (follow): Both knots are in the same position. No need to follow")
		return
	}
	if dist(knot, followedKnot) <= 1 {
		log.Println("INFO (follow): Knots are at distance 1 or less. No need to follow")
		return
	}
	// Moving knot
	// Place the knot in the same position as the knot it wants to follow
	// Move one position in the opposite direction, so it ends behind it
	knot.x = followedKnot.x
	knot.y = followedKnot.y
	knot.move(opposite[followedKnot.direction])

}

func main() {
	scanner, f := read("input.txt")
	defer f.Close()
	head := Knot{x: 0, y: 0}
	tail := Knot{x: 0, y: 0}
	visited := []string{}

	for scanner.Scan() {
		direction := strings.Split(scanner.Text(), " ")[0]
		steps, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[1])

		// Avoid direction to be empty
		if tail.direction == "" {
			tail.direction = direction
		}

		if head.direction == "" {
			head.direction = direction
		}
		for i := 0; i < steps; i++ {
			head.move(direction)
			tail.follow(head)
			log.Println("INFO: Head (", head.x, head.y, head.direction, ")", "Tail (", tail.x, tail.y, tail.direction, ")")
			coord := "(" + fmt.Sprint(tail.x) + "," + fmt.Sprint(tail.y) + ")"
			if !contains(visited, coord) {
				visited = append(visited, coord)
			}
		}
	}
	fmt.Println("Final state: Head (", head.x, head.y, head.direction, ")", "Tail (", tail.x, tail.y, tail.direction, ")")
	fmt.Println("Visited: ", len(visited))
	//fmt.Println(visited)
}
