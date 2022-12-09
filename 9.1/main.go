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

	headPos := Position{x: 1, y: 5}
	tailPos := Position{x: 1, y: 5}
	var tailPositions []string

	printGrid(&headPos, &tailPos)

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, " ")

		dir := tmp[0]
		length, _ := strconv.Atoi(tmp[1])

		moveHeadBy(&headPos, &tailPos, dir, length, &tailPositions)
	}

	fmt.Println(len(tailPositions))
}

func moveHeadBy(head *Position, tail *Position, dir string, length int, tailPositions *[]string) {
	fmt.Println(dir, " ", length)

	for i := 1; i <= length; i++ {
		fmt.Println("HEAD from ", head)

		switch dir {
		case "U":
			head.y = head.y - 1
		case "D":
			head.y = head.y + 1
		case "R":
			head.x = head.x + 1
		case "L":
			head.x = head.x - 1
		}

		// fmt.Println("HEAD to", head)
		// fmt.Println("TAIL from ", tail)

		nextTailPosition := findNextTailPosition(head, tail)

		fmt.Println("TAIL to", nextTailPosition)

		tail.x = nextTailPosition.x
		tail.y = nextTailPosition.y

		posId := strconv.Itoa(tail.x) + "-" + strconv.Itoa(tail.y)

		fmt.Println(posId)

		if !contains(*tailPositions, posId) {
			*tailPositions = append(*tailPositions, posId)
		}

		printGrid(head, tail)

		fmt.Println("=====================================================================")
	}

	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
}

func printGrid(head *Position, tail *Position) {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 6; j++ {
			printed := false

			if i == head.y && j == head.x {
				fmt.Print("H")
				printed = true
			}

			if !printed {
				if i == tail.y && j == tail.x {
					fmt.Print("T")
					printed = true
				}
			}

			if !printed {
				fmt.Print(".")
			}
		}

		fmt.Println("")
	}
}

func findNextTailPosition(head *Position, tail *Position) Position {
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
		fmt.Println("down")
		return 1
	} else {
		fmt.Println("up")
		return -1
	}
}

func moveHorizotally(head Position, tail Position) int {
	if head.x > tail.x {
		fmt.Println("right")
		return 1
	} else {
		fmt.Println("left")
		return -1
	}
}
