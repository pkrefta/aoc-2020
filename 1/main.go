package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentElfCalories := 0

	var best [3]int

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			calories, err := strconv.Atoi(line)

			if err == nil {
				currentElfCalories += calories
			}
		} else {
			for i := 0; i < 3; i++ {
				if currentElfCalories > best[i] {
					if i == 0 {
						best[2] = best[1]
						best[1] = best[0]
					}

					if i == 1 {
						best[2] = best[1]
					}

					best[i] = currentElfCalories

					break
				}
			}

			currentElfCalories = 0
		}
	}

	result := 0

	for _, v := range best {
		result += v
	}

	fmt.Println(result)
}
