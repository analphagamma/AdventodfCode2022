package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// open file
	f, _ := os.Open("day04.txt")
	defer f.Close()

	// slice to hold file data
	rucksacks := make([]string, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, line)
	}
	var sum int
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, sack := range rucksacks {
		// splitting it into two compartments
		// and finding common element
		var letter string
		for _, comp1 := range sack[:len(sack)/2] {
			for _, comp2 := range sack[len(sack)/2:] {
				if comp1 == comp2 {
					letter = string(comp1)
					continue
				}
			}
		}
		// find priority value and add it to sum
		sum += strings.Index(alphabet, letter) + 1
	}
	fmt.Println("Sum of priorities: ", sum)
}
