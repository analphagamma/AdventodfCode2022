package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// open file
	f, _ := os.Open("day01.txt")
	defer f.Close()

	// slice to hold file data
	cals := make([]int, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt, _ := strconv.Atoi(line)
		cals = append(cals, lineInt)
	}

	// slice to hold calorie values of individual elves
	elves := make([]int, 0)
	// id of elf :)
	var elfindex int
	// first elf
	elves = append(elves, 0)
	for _, cal := range cals {
		if cal == 0 {
			elfindex++
			elves = append(elves, 0)
			continue
		} else {
			elves[elfindex] += cal
		}
	}
	// sort slice in descending order
	sort.Slice(elves, func(a, b int) bool {
		return elves[b] < elves[a]
	})
	fmt.Println("Largest 3 values: ", elves[0]+elves[1]+elves[2])
}
