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
	rules := map[string]string{
		"A": "Y", // Rock is beaten by Paper
		"B": "Z", // Paper is beaten by Scissors
		"C": "X", // Scissors is beaten by Paper
	}
	scores := map[string]int{
		"A": 1, // for comparison
		"B": 2,
		"C": 3,
		"X": 1, // actual points
		"Y": 2,
		"Z": 3,
	}
	var total_points int
	for _, game := range games {
		if rules[game[0]] == game[1] { // you win
			total_points += scores[game[1]] + 6
		} else if scores[game[0]] == scores[game[1]] { // it's a draw
			total_points += scores[game[1]] + 3
		} else {
			total_points += scores[game[1]] // AI wins
		}
	}
	fmt.Printf("Total points: %d\n", total_points)
}
