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
	var l1 []string
	var l2 []string
	for scanner.Scan() {
		sp := strings.SplitAfter(scanner.Text(), "")
		l1, l2 = sp[:(len(sp)/2)], sp[(len(sp)/2):]
		for _, l := range l1 {
			if contains(l2, l) {
				id := strings.ToLower(l)
				sum += upperMultiplier(l) + points[id]
				fmt.Printf("Lista1: %s, Lista2: %s. Contiene: %s , id: %s, puntos: %d , multi: %d, total: %d\n", l1, l2, l, id, points[id], upperMultiplier(l), upperMultiplier(l)+points[id])
				break
			}
		}

	}
	fmt.Print(points)
	fmt.Print(sum)
}
