package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sectionsToRange(sectionsStr string) []int {
	tmp := strings.Split(sectionsStr, "-")

	var out []int

	start, _ := strconv.Atoi(tmp[0])
	end, _ := strconv.Atoi(tmp[1])

	for i := start; i <= end; i++ {
		out = append(out, i)
	}

	return out
}

func sectionContainsSectionId(id int, ids []int) bool {
	for _, el := range ids {
		if el == id {
			return true
		}
	}

	return false
}

func sectionsContainSections(firstSections []int, secondSections []int) bool {
	firstLen := len(firstSections)
	secondLen := len(secondSections)

	fmt.Println(firstSections)
	fmt.Println(secondSections)

	if firstLen <= secondLen {
		fmt.Println("first check")
		allIn := true
		// first is "in" second
		for _, sectionId := range firstSections {
			if !sectionContainsSectionId(sectionId, secondSections) {
				allIn = false
				break
			}
		}

		if allIn {
			return true
		}
	}

	if secondLen <= firstLen {
		fmt.Println("second check")
		// second is "in" first
		allIn := true
		// first is "in" second
		for _, sectionId := range secondSections {
			if !sectionContainsSectionId(sectionId, firstSections) {
				allIn = false
				break
			}
		}

		if allIn {
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

	overlapping := 0

	for scanner.Scan() {
		line := scanner.Text()

		tmp := strings.Split(line, ",")

		firstSections := sectionsToRange(tmp[0])
		secondSections := sectionsToRange(tmp[1])
	}
}
