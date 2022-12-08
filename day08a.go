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
	f, _ := os.Open("day08.txt")
	defer f.Close()

	// slice to hold file data
	heatMap := make([][]int, 0)
	// read file into slice
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, num := range strings.Split(line, "") {
			num, _ := strconv.Atoi(num)
			row = append(row, num)
		}
		heatMap = append(heatMap, row)
	}

	isVisible := func(tree int, neighbours []int) bool {
		for _, n := range neighbours {
			if n >= tree {
				return false
			}
		}
		return true
	}
	var visibleTrees int
	// go through rows
	for i, row := range heatMap {
		// got through individual trees
		for j, tree := range row {

			if i == 0 || j == 0 || i == len(heatMap)-1 || j == len(row)-1 { //edge of the map
				visibleTrees++
			} else {
				// left side
				var leftNeighbours []int
				for k := 0; k < j; k++ {
					leftNeighbours = append(leftNeighbours, heatMap[i][k])
				}
				if isVisible(tree, leftNeighbours) {
					visibleTrees++
					continue
				}
				// right side
				var rightNeighbours []int
				for k := j + 1; k < len(row); k++ {
					rightNeighbours = append(rightNeighbours, heatMap[i][k])
				}
				if isVisible(tree, rightNeighbours) {
					visibleTrees++
					continue
				}
				// top
				var topNeighbours []int
				for k := 0; k < i; k++ {
					topNeighbours = append(topNeighbours, heatMap[k][j])
				}
				if isVisible(tree, topNeighbours) {
					visibleTrees++
					continue
				}
				// bottom
				var bottomNeighbours []int
				for k := i + 1; k < len(heatMap); k++ {
					bottomNeighbours = append(bottomNeighbours, heatMap[k][j])
				}
				if isVisible(tree, bottomNeighbours) {
					visibleTrees++
					continue
				}
			}
		}
	}

	// display results
	fmt.Println("Visible trees: ", visibleTrees)
}
