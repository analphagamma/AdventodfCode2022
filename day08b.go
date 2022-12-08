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
	var scenicMax int
	// go through rows
	for i, row := range heatMap {
		// got through individual trees
		for j, tree := range row {
			if i == 0 || j == 0 || i == len(heatMap)-1 || j == len(row)-1 { //edge of the map
				continue
			}
			scenicScore := 1
			var visibleTrees int
			// left side
			for k := j - 1; k >= 0; k-- {
				if heatMap[i][k] >= tree {
					visibleTrees++
					break
				}
				visibleTrees++
			}
			scenicScore *= visibleTrees
			visibleTrees = 0
			// right side
			for k := j + 1; k < len(row); k++ {
				if heatMap[i][k] >= tree {
					visibleTrees++
					break
				}
				visibleTrees++
			}
			scenicScore *= visibleTrees
			visibleTrees = 0
			// top
			for k := i - 1; k >= 0; k-- {
				if heatMap[k][j] >= tree {
					visibleTrees++
					break
				}
				visibleTrees++
			}
			scenicScore *= visibleTrees
			visibleTrees = 0
			// bottom
			for k := i + 1; k < len(heatMap); k++ {
				if heatMap[k][j] >= tree {
					visibleTrees++
					break
				}
				visibleTrees++
			}
			scenicScore *= visibleTrees
			visibleTrees = 0

			if scenicScore > scenicMax {
				scenicMax = scenicScore
			}
		}
	}

	// display results
	fmt.Println("Highest scenic score: ", scenicMax)
}
