package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func removeIndex(s []string, index int) []string {
	t := make([]string, len(s))
	copy(t, s)
	return append(t[:index], t[index+1:]...)
}

func uniq(l []string) bool {
	for i, v := range l {
		tmp := removeIndex(l, i)
		fmt.Printf("INFO: marker: %q, index: %d, tmp: %q, v: %s\n", l, i, tmp, v)
		if contains(tmp, v) {
			return false
		}
	}

	return true
}

func main() {
	scanner, f := read()
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		l := strings.Split(line, "")
		var marker []string
		for i, char := range l {
			marker = append(marker, char)

			if len(marker) == 15 {
				marker = marker[1:]
			}
			fmt.Println(marker)
			if len(marker) == 14 {
				if uniq(marker) {
					fmt.Println(i+1, ": ", marker)
					break
				}
			}
		}
	}
}
