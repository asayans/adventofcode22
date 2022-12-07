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

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func contains(elems []int, v int) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	scanner, f := read()
	defer f.Close()

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		p1, p2 := strings.Split(line, ",")[0], strings.Split(line, ",")[1]
		p1_min, _ := strconv.Atoi(strings.Split(p1, "-")[0])
		p1_max, _ := strconv.Atoi(strings.Split(p1, "-")[1])
		p2_min, _ := strconv.Atoi(strings.Split(p2, "-")[0])
		p2_max, _ := strconv.Atoi(strings.Split(p2, "-")[1])

		p1_range := makeRange(p1_min, p1_max)
		p2_range := makeRange(p2_min, p2_max)

		for _, n := range p1_range {
			if contains(p2_range, n) {
				res++
				break
			}
		}
	}

	fmt.Print(res)
}
