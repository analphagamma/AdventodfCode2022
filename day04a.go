package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, _ := os.Open("day04.txt")
	defer f.Close()

	// slice to hold file data
	assigments := make([]string, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		assigments = append(assigments, line)
	}
	//assigments := [6]string{
	//	"2-4,6-8",
	//	"2-3,4-5",
	//	"5-7,7-9",
	//	"2-8,3-7",
	//	"6-6,4-6",
	//	"2-6,4-8",
	//}
	var overlaps int
	for _, line := range assigments {
		assigmentPair := strings.Split(line, ",")
		elf1, elf2 := assigmentPair[0], assigmentPair[1]
		var elf1ids = getBoundaries(elf1)
		var elf2ids = getBoundaries(elf2)

		if (elf1ids[0] <= elf2ids[0] && elf1ids[1] >= elf2ids[1]) || (elf2ids[0] <= elf1ids[0] && elf2ids[1] >= elf1ids[1]) {
			overlaps++
		}
	}
	fmt.Println("Number of overlaps: ", overlaps)
}

func getBoundaries(rangeStr string) [2]int {
	idRange := strings.Split(rangeStr, "-")
	r_low, _ := strconv.Atoi(idRange[0])
	r_high, _ := strconv.Atoi(idRange[1])

	return [2]int{r_low, r_high}

}
