package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("full_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var stackLines []string
	var numbers []string
	var stacks = make(map[string][]string)

	readingStacks := true

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			numbers = strings.Split(strings.Trim(stackLines[len(stackLines)-1], " "), "   ")
			readStacks(stacks, stackLines, numbers)
			readingStacks = false
		} else {

			if readingStacks {
				stackLines = append(stackLines, line)
			} else {
				applyMove(stacks, line)
			}
		}
	}

	for _, number := range numbers {
		stack := stacks[number]
		fmt.Print(stack[len(stack)-1])
	}
}

func applyMove(stacks map[string][]string, moveLine string) {
	tmp := strings.Split(moveLine, " ")
	amount, _ := strconv.Atoi(tmp[1])

	from := tmp[3]
	to := tmp[5]
	fromLen := len(stacks[from])

	toMove := stacks[from][fromLen-amount : fromLen]

	stacks[from] = stacks[from][:fromLen-amount]
	stacks[to] = append(stacks[to], toMove...)
}

func readStacks(stacks map[string][]string, stackLines []string, numbers []string) {
	for _, number := range numbers {
		stacks[number] = nil
	}

	for i := len(stackLines) - 1; i >= 0; i-- {
		for j, number := range numbers {
			line := " " + stackLines[i]

			chunk := line[j*4 : (j+1)*4]

			if strings.Contains(chunk, " [") {
				letter := strings.TrimSuffix(strings.TrimPrefix(chunk, " ["), "]")

				stacks[number] = append(stacks[number], letter)
			}
		}
	}
}
