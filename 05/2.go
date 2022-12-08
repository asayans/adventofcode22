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

func reverseSliceString(s []string) []string {
	rev := []string{}
	for i := range s {
		rev = append(rev, s[len(s)-1-i])
	}
	return rev
}

func move(ship map[int][]string, qty int, srci int, dsti int) map[int][]string {
	src := ship[srci]
	dst := ship[dsti]
	toMove := src[len(src)-qty:]
	//fmt.Printf("INFO: \n\t ship: %q \n\t src: %q \n\t dst: %q \n\t qty: %d , \n\t toMove: %q \n", ship, src, dst, qty, toMove)
	for _, crate := range toMove {
		dst = append(dst, crate)
	}
	src = src[:len(src)-qty]
	ship[srci], ship[dsti] = src, dst
	return ship
}

func createShip(data [][]string) map[int][]string {
	ship := make(map[int][]string)
	for _, stack := range data {
		//fmt.Printf("\n Line: %d \n", x)
		for i, crate := range stack {
			//fmt.Printf("x: %d, i: %d, data: %s, crate: %s \n", x, i, stack, crate)
			if crate != "_" {
				ship[i+1] = append([]string{crate}, ship[i+1]...)
			}
		}
	}
	return ship
}

func showShip(ship map[int][]string) {
	for k, v := range ship {
		fmt.Println(k, v)
	}
}
func main() {
	var data [][]string
	scanner, f := read()
	defer f.Close()
	shipCreated := false
	var moves [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if shipCreated {
			// Reading move actions
			qty, _ := strconv.Atoi(strings.Split(line, " ")[1])
			src, _ := strconv.Atoi(strings.Split(line, " ")[3])
			dst, _ := strconv.Atoi(strings.Split(line, " ")[5])
			//fmt.Printf("qty: %d, src: %d, dst: %d \n", qty, src, dst)
			moves = append(moves, []int{qty, src, dst})
			//fmt.Printf("moves: %d", moves)
		} else if line != "" && !shipCreated {
			line = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(line, "[", ""), "]", ""), "    ", " _")
			data = append(data, strings.Split(line, " "))
			continue
		} else {
			shipCreated = true
		}
	}

	ship := createShip(data[:len(data)-1])
	//fmt.Print(ship)

	for _, m := range moves {
		qty := m[0]
		src := m[1]
		dst := m[2]
		//fmt.Printf("\nMove %d: \nMove %d crates from %d to %d \n", i, qty, src, dst)
		//showShip(ship)
		ship = move(ship, qty, src, dst)
	}
	fmt.Println("Final ship state:")
	showShip(ship)
}
