package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Read file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	actual := 0
	total_list := []int{}

	for scanner.Scan() {
		numStr := scanner.Text()

		if numStr != "" {
			num, errStr := strconv.Atoi(numStr)

			if errStr != nil {
				log.Fatal(errStr)
			}

			actual += num

		} else {
			total_list = append(total_list, actual)
			actual = 0
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(total_list)))

	fmt.Print(total_list[0:3])
	fmt.Print(total_list[0] + total_list[1] + total_list[2])
}
