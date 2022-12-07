package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func read() (*bufio.Scanner, *os.File) {
	f, err := os.Open("input.txt")
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

func upperMultiplier(s string) int {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return 26
		}
	}
	return 0
}

func main() {
	scanner, f := read()
	defer f.Close()
	points := make(map[string]int)

	for i, r := range "abcdefghijklmnopqrstuvwxyz" {
		letter := string(r)
		points[letter] = i + 1
	}
	sum := 0
	cont := 0
	var rucks [][]string
	for scanner.Scan() {
		line := strings.SplitAfter(scanner.Text(), "")
		rucks = append(rucks, line)

		if cont == 2 {
			for _, l := range rucks[0] {
				if contains(rucks[1], l) {
					if contains(rucks[2], l) {
						id := strings.ToLower(l)
						sum += upperMultiplier(l) + points[id]
						fmt.Printf("Ruck1: %s \nRuck2: %s \nRuck3: %s \nContiene: %s , id: %s, puntos: %d , multi: %d, total: %d\n\n", rucks[0], rucks[1], rucks[2], l, id, points[id], upperMultiplier(l), upperMultiplier(l)+points[id])
						break
					}
				}
			}
			cont = 0
			rucks = nil
		} else {
			cont++
		}

	}
	fmt.Print(sum)
}
