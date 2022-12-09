package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("full_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var grid [][]int

	scanner := bufio.NewScanner(f)
	lineNumber := 0

	for scanner.Scan() {
		grid = append(grid, nil)

		line := scanner.Text()

		for _, h := range line {
			intHeight, _ := strconv.Atoi(string(h))
			grid[lineNumber] = append(grid[lineNumber], intHeight)
		}

		lineNumber++
	}

	visibleCount := 0
	highestScore := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {

			visible := isVisible(grid, row, col)
			score := calculateScenicScore(grid, row, col)

			if score > highestScore {
				highestScore = score
			}

			if visible {
				visibleCount++
			}
		}
	}

	fmt.Println(visibleCount)
	fmt.Println(highestScore)
}

func calculateScenicScore(grid [][]int, row int, col int) int {
	size := len(grid[0])
	treeHeight := grid[row][col]

	distLeft, distRight, distTop, distBottom := 0, 0, 0, 0

	// to the left
	for i := col - 1; i >= 0; i-- {
		distLeft++
		otherHeight := grid[row][i]

		if otherHeight >= treeHeight {
			break
		}
	}

	// to the top

	for i := row - 1; i >= 0; i-- {
		distTop++
		otherHeight := grid[i][col]

		if otherHeight >= treeHeight {
			break
		}

	}

	// to the right

	for i := col + 1; i < size; i++ {
		distRight++
		otherHeight := grid[row][i]

		if otherHeight >= treeHeight {
			break
		}

	}

	// to the bottom

	for i := row + 1; i < size; i++ {
		distBottom++
		otherHeight := grid[i][col]

		if otherHeight >= treeHeight {
			break
		}

	}

	return distTop * distBottom * distLeft * distRight
}

func isVisible(grid [][]int, row int, col int) bool {
	size := len(grid[0])
	treeHeight := grid[row][col]
	visible := true

	if row == 0 || col == 0 || row == size-1 || col == size-1 {
		return visible
	}

	// to the left

	for i := col - 1; i >= 0; i-- {
		otherHeight := grid[row][i]

		if otherHeight >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	} else {
		visible = true
	}

	// to the top

	for i := row - 1; i >= 0; i-- {
		otherHeight := grid[i][col]

		if otherHeight >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	} else {
		visible = true
	}

	// to the right

	for i := col + 1; i < size; i++ {
		otherHeight := grid[row][i]

		if otherHeight >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	} else {
		visible = true
	}

	// to the bottom

	for i := row + 1; i < size; i++ {
		otherHeight := grid[i][col]

		if otherHeight >= treeHeight {
			visible = false
			break
		}
	}

	if visible {
		return true
	}

	return false
}
