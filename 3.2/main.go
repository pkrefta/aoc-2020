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

	for r, i := 'a', 1; r <= 'z'; r, i = r+1, i+1 {
		priorities[r] = i
		R := unicode.ToUpper(r)
		priorities[R] = i + 26
	}

	scanner := bufio.NewScanner(f)
	total := 0

	var rucksacks []string

	for scanner.Scan() {
		rucksack := scanner.Text()

		rucksacks = append(rucksacks, rucksack)

		if len(rucksacks) == 3 {
			for c, priority := range priorities {

				if strings.ContainsRune(rucksacks[0], c) && strings.ContainsRune(rucksacks[1], c) && strings.ContainsRune(rucksacks[2], c) {
					fmt.Println(string(c))
					total += priority
				}
			}

			rucksacks = nil
		}
	}

	fmt.Println(total)
}
