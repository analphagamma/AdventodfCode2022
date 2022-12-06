package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// open file
	f, _ := os.ReadFile("day06.txt")

	// slice to hold file data
	signal := string(f)

	var firstMarker int
	for i := 0; i < len(signal); i++ {
		markers := make(map[string]int)
		for _, char := range strings.Split(signal[i:i+14], "") {
			markers[char] += 1
		}
		if len(markers) == 14 {
			firstMarker = i + 14
			break
		}
		if firstMarker != 0 {
			break
		}
	}
	fmt.Printf("First marker: %d\n", firstMarker)
}
