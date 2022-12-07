package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func rps(p1 string, p2 string) int {

	ref := make(map[string]string)
	ref["X"] = "A"
	ref["Y"] = "B"
	ref["Z"] = "C"

	res := 0
	draw := p1 == ref[p2]

	if draw {
		res = 3
	} else {
		switch p1 {
		case "A":
			if p2 == "Y" {
				res = 6
			}
		case "B":
			if p2 == "Z" {
				res = 6
			}
		case "C":
			if p2 == "X" {
				res = 6
			}
		}
	}

	switch p2 {
	case "X":
		res += 1
	case "Y":
		res += 2
	case "Z":
		res += 3
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
	var p1, p2 string

	for scanner.Scan() {
		line = scanner.Text()
		p1, p2 = strings.Fields(line)[0], strings.Fields(line)[1]
		// fmt.Print(p1, p2+"\n")
		round_score := rps(p1, p2)
		score += round_score
	}

	fmt.Print(score)
	//fmt.Print(rps("C", "Z"))
}
