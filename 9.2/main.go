package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func contains(haystack []string, needle string) bool {
	for _, r := range haystack {
		if r == needle {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("full_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	knots := make([]Position, 10)

	for i := range knots {
		knots[i].x = 1
		knots[i].y = 5
	}

	printGrid(knots)

	var tailPositions []string

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")

		dir := tmp[0]
		length, _ := strconv.Atoi(tmp[1])

		moveHeadBy(knots, dir, length, &tailPositions)
	}

	fmt.Println(len(tailPositions))
}

func moveHeadBy(knots []Position, dir string, length int, tailPositions *[]string) {
	for i := 1; i <= length; i++ {
		head := knots[0]

		switch dir {
		case "U":
			knots[0].y = head.y - 1
		case "D":
			knots[0].y = head.y + 1
		case "R":
			knots[0].x = head.x + 1
		case "L":
			knots[0].x = head.x - 1
		}

		for idx := range knots {
			nextIdx := idx + 1

			if nextIdx < len(knots) {
				nextKnotPosition := findNextKnotPosition(&knots[idx], &knots[nextIdx])

				knots[nextIdx].x = nextKnotPosition.x
				knots[nextIdx].y = nextKnotPosition.y

				if nextIdx == 9 {
					tail := knots[nextIdx]
					posId := strconv.Itoa(tail.x) + "-" + strconv.Itoa(tail.y)

					if !contains(*tailPositions, posId) {
						*tailPositions = append(*tailPositions, posId)
					}
				}
			}
		}

		printGrid(knots)
	}
}

func printGrid(knots []Position) {
	return

	for i := -20; i <= 20; i++ {
		for j := -20; j <= 30; j++ {
			printed := false

			for idx := range knots {
				knot := knots[idx]

				if i == knot.y && j == knot.x {
					if idx == 0 {
						fmt.Print("H")
					} else {
						fmt.Print(idx)
					}

					printed = true
					break
				}
			}

			if !printed {
				fmt.Print(".")
			}
		}

		fmt.Println("")
	}
}

func findNextKnotPosition(head *Position, tail *Position) Position {
	x, y := tail.x, tail.y

	xDist := math.Abs(float64(head.x - tail.x))
	yDist := math.Abs(float64(head.y - tail.y))

	if xDist < 2 && yDist < 2 {
		// we don't have to move
		return *tail
	} else {
		if yDist == 0 {
			x = x + moveHorizotally(*head, *tail)
		} else if xDist == 0 {
			y = y + moveVertically(*head, *tail)
		} else {
			x = x + moveHorizotally(*head, *tail)
			y = y + moveVertically(*head, *tail)
		}
	}

	return Position{x: x, y: y}
}

func moveVertically(head Position, tail Position) int {
	if head.y > tail.y {
		return 1
	} else {
		return -1
	}
}

func moveHorizotally(head Position, tail Position) int {
	if head.x > tail.x {
		return 1
	} else {
		return -1
	}
}
