package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// open file
	f, _ := os.Open("day03.txt")
	defer f.Close()

	// slice to hold file data
	games := make([][]string, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		games = append(games, strings.Split(line, " "))
	}

	beats := map[string]string{
		"A": "B", // Rock is beaten by Paper
		"B": "C", // Paper is beaten by Scissors
		"C": "A", // Scissors is beaten by Paper
	}
	scores := map[string]int{
		"A": 1, // for comparison
		"B": 2,
		"C": 3,
	}
	var total_points int
	for _, game := range games {
		fmt.Println(game)
		switch game[1] {
		case "X": // AI wins
			for k, v := range beats {
				if v == game[0] {
					total_points += scores[k]
				}
			}
		case "Y": // Draw
			total_points += scores[game[0]] + 3
		case "Z": // Player wins
			total_points += scores[beats[game[0]]] + 6
		}
	}
	fmt.Printf("Total points: %d\n", total_points)
}
