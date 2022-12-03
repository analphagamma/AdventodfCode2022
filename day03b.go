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
	group := make([]string, 0)
	for _, sack := range rucksacks {
		// add sack to group
		if len(group) < 3 {
			group = append(group, sack)
		}
		if len(group) == 3 {
			// iterating through the three sacks
			var letter string
			for _, let1 := range group[0] {
				for _, let2 := range group[1] {
					if let1 == let2 {
						for _, let3 := range group[2] {
							if let2 == let3 {
								letter = string(let1)
								continue
							}
						}
					}
				}
			}
			// find priority value and add it to sum
			sum += strings.Index(alphabet, letter) + 1
			// empty group slice
			group = nil
		} else {
			continue
		}
	}
	fmt.Println("Sum of priorities: ", sum)
}
