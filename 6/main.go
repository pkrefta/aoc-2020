package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func indexOf(haystack []string, needle string) int {
	for idx, r := range haystack {
		if r == needle {
			return idx
		}
	}
	return -1
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	f, err := os.Open("full_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

SCANNER:
	for scanner.Scan() {
		line := scanner.Text()

		var tmp []string

		for i, c := range line {
			idx := indexOf(tmp, string(c))

			if idx != -1 {
				tmp = tmp[idx+1:]
			}

			tmp = append(tmp, string(c))

			if len(tmp) == 14 {
				marker := i + 1
				fmt.Println(marker)
				continue SCANNER
			}

		}
	}

}
