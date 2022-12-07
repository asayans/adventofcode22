package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func rps(p1 string, p2 string) int {
	res := 0
	draw := p1 == p2

	if draw {
		res = 3
	} else {
		switch p1 {
		case "A":
			if p2 == "B" {
				res = 6
			}
		case "B":
			if p2 == "C" {
				res = 6
			}
		case "C":
			if p2 == "A" {
				res = 6
			}
		}
	}

	switch p2 {
	case "A":
		res += 1
	case "B":
		res += 2
	case "C":
		res += 3
	}

	return res
}

func chooseMove(p1 string, e2 string) string {
	win := make(map[string]string)
	lose := make(map[string]string)

	var res string

	win["A"] = "B"
	win["B"] = "C"
	win["C"] = "A"

	lose["A"] = "C"
	lose["B"] = "A"
	lose["C"] = "B"

	switch e2 {
	case "X":
		res = lose[p1]
	case "Y":
		res = p1
	case "Z":
		res = win[p1]
	}

	return res
}

func main() {

	// Read file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	score := 0
	var line string
	var p1, p2, e2 string

	for scanner.Scan() {
		line = scanner.Text()
		p1, e2 = strings.Fields(line)[0], strings.Fields(line)[1]
		p2 = chooseMove(p1, e2)
		// fmt.Print(p1, p2+"\n")
		round_score := rps(p1, p2)
		score += round_score
	}

	fmt.Print(score)
	//fmt.Print(rps("C", "Z"))
}
