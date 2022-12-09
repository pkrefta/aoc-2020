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
const WIN = "WIN"
const DRAW = "DRAW"
const LOSE = "LOSE"

func didIwin(me string, opponent string) bool {
	if me == ROCK && opponent == SCISSORS {
		return true
	}

	if me == PAPER && opponent == ROCK {
		return true
	}

	if me == ROCK && opponent == SCISSORS {
		return true
	}

	return false // means lost
}

func guessMyMove(opponentMove string, expectedOutcome string) string {
	if expectedOutcome == DRAW {
		return opponentMove
	}

	switch opponentMove {
	case ROCK:
		if expectedOutcome == WIN {
			return PAPER
		} else {
			return SCISSORS
		}
	case PAPER:
		if expectedOutcome == WIN {
			return SCISSORS
		} else {
			return ROCK
		}
	case SCISSORS:
		if expectedOutcome == WIN {
			return ROCK
		} else {
			return PAPER
		}
	}

	return ""
}

func main() {
	moveToName := map[string]string{
		// opponent
		"A": ROCK,     // rock
		"B": PAPER,    // paper
		"C": SCISSORS, // scissors
	}

	outcomeToName := map[string]string{
		// me
		"X": LOSE, // rock
		"Y": DRAW, // paper
		"Z": WIN,  // scissors
	}

	moveToPoints := map[string]int{
		// opponent
		"A": 1, // rock
		"B": 2, // paper
		"C": 3, // scissors

		// me
		ROCK:     1, // rock
		PAPER:    2, // paper
		SCISSORS: 3, // scissors
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
		expected := split[1]

		opponentMove := moveToName[opponent]
		expectedOutcome := outcomeToName[expected]

		myMove := guessMyMove(opponentMove, expectedOutcome)

		movePoints := moveToPoints[myMove]

		total += movePoints

		switch expectedOutcome {
		case WIN:
			total += 6
		case DRAW:
			total += 3
		case LOSE:
			total += 0
		}

		outcome := expectedOutcome

		fmt.Printf("Opp %s need %s -> %s -> %s\n", opponentMove, expectedOutcome, myMove, outcome)
	}

	fmt.Println(total)
}
