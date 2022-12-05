package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, _ := os.Open("day05.txt")
	defer f.Close()

	// slice to hold file data
	instructions := make([]string, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}
	// ok, I'm not gonna parse this
	crates := [][]string{
		{}, // indexes start at 0 but the crate numbers start with 1
		{"Q", "S", "W", "C", "Z", "V", "F", "T"},
		{"Q", "R", "B"},
		{"B", "Z", "T", "Q", "P", "M", "S"},
		{"D", "V", "F", "R", "Q", "H"},
		{"J", "G", "L", "D", "B", "S", "T", "P"},
		{"W", "R", "T", "Z"},
		{"H", "Q", "M", "N", "S", "F", "R", "J"},
		{"R", "N", "F", "H", "W"},
		{"J", "Z", "T", "Q", "P", "R", "B"},
	}

	re := regexp.MustCompile(`\d+`)
	for _, inst := range instructions {
		moves := re.FindAllString(inst, -1)
		move := make([]int, 3)
		for i, elem := range moves {
			move[i], _ = strconv.Atoi(elem)
		}
		for i := 1; i <= move[0]; i++ {
			// we store the crate we're moving
			var crate string
			// pick stack we're modifying
			stack := crates[move[1]]
			// pop last element of stack indicated in move[1]
			// store it in "crate" and modify original crate
			crate, crates[move[1]] = stack[len(stack)-1], stack[:len(stack)-1]
			// add 'crate' to destination
			crates[move[2]] = append(crates[move[2]], crate)
		}
	}
	topBoxes := make([]string, len(crates))
	for i, crate := range crates {
		if len(crate) > 0 {
			topBoxes[i] = crate[len(crate)-1]
		}
	}
	fmt.Printf("Top boxes: %s\n", strings.Join(topBoxes, ""))
}
