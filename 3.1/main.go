package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("full_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	priorities := make(map[rune]int)

	for r, i := 'a', 1; r < 'z'; r, i = r+1, i+1 {
		priorities[r] = i
		R := unicode.ToUpper(r)
		priorities[R] = i + 26
	}

	scanner := bufio.NewScanner(f)
	total := 0

	for scanner.Scan() {
		rucksack := scanner.Text()

		middle := len(rucksack) / 2

		firstCompartment := rucksack[0:middle]
		secondCompartment := rucksack[middle:]

		for char, priority := range priorities {
			if strings.ContainsRune(firstCompartment, char) && strings.ContainsRune(secondCompartment, char) {
				total += priority
			}
		}
	}

	fmt.Println(total)
}
