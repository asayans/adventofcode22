package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res := 0
	cont := 0
	top := 0
	actual := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numStr := scanner.Text()

		if numStr != "" {
			num, errStr := strconv.Atoi(numStr)

			if errStr != nil {
				log.Fatal(errStr)
			}

			actual += num

		} else {
			if actual > top {
				res = cont
				top = actual
			}
			cont += 1
			actual = 0
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(res, top)
}
