package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func move(port [][]string, qty int, srci int, dsti int) [][]string {
	src := port[srci]
	dst := port[dsti]
	for _, crate := range src[len(src)-qty:] {
		dst = append(dst, crate)
	}
	src = src[len(src)-qty:]
	port[srci], port[dsti] = src, dst
	return port
}

func read() (*bufio.Scanner, *os.File) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	return scanner, f
}

func main() {
	var port [][]string

	scanner, f := read()
	defer f.Close()

	for scanner.Scan() {
		//line := strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), "[", ""), "]", "")
		line := scanner.Text()
		if line == "" {
			port = port[:len(port)-1]
			fmt.Print(port[0][0], cap(port[0]))
			break
		}
		l := []string{line}
		port = append(port, l)
	}
}
