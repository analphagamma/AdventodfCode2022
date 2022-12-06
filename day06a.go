package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// open file
	f, _ := os.ReadFile("day06.txt")

	// var to hold file data
	signal := string(f)

	var firstMarker int
	for i := 0; i < len(signal); i++ {
		markers := make(map[string]int)
		for _, char := range strings.Split(signal[i:i+4], "") {
			markers[char] += 1
		}
		if len(markers) == 4 {
			firstMarker = i + 4
			break
		}
	}
	fmt.Printf("First marker: %d\n", firstMarker)
}
