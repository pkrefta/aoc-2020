package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const ROCK = "ROCK"
const PAPER = "PAPER"
const SCISSORS = "SCISSORS"

func didIwin(me string, opponent string) bool {
	if me == ROCK {
		if opponent == SCISSORS {
			return true
		}
	}

	if me == PAPER {
		if opponent == ROCK {
			return true
		}
	}

	if me == ROCK {
		if opponent == SCISSORS {
			return true
		}
	}

	return false // means lost
}

func main() {
	moveToName := map[string]string{
		// opponent
		"A": ROCK,     // rock
		"B": PAPER,    // paper
		"C": SCISSORS, // scissors

		// me
		"X": ROCK,     // rock
		"Y": PAPER,    // paper
		"Z": SCISSORS, // scissors
	}

	moveToPoints := map[string]int{
		// opponent
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors

		// me
		"X": 1, // rock
		"Y": 2, // paper
		"Z": 3, // scissors
	}

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")

		opponent := split[0]
		me := split[1]

		myMove := moveToName[me]
		opponentMove := moveToName[opponent]

		total += moveToPoints[me]

		outcome := ""

		if myMove == opponentMove {
			outcome = "DRAW"
			total += 3
		} else {
			if didIwin(myMove, opponentMove) {
				total += 6
				outcome = "WIN"
			} else {
				outcome = "LOSE"
			}
		}

		fmt.Printf("Opp %s ME %s -> %s\n", opponentMove, myMove, outcome)
	}

	fmt.Println(total)
}
