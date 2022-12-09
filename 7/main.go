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

	var lines []string

	contents := make(map[string][]string)
	sizes := make(map[string]int)

	var currentPathItems []string
	currentPath := ""

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		tmp := strings.Split(line, " ")

		if tmp[0] == "$" {
			if tmp[1] == "cd" {
				if tmp[2] != ".." {
					currentPathItems = append(currentPathItems, tmp[2])
				} else {
					currentPathItems = currentPathItems[:len(currentPathItems)-1]
				}
			}

			if tmp[1] == "ls" {
				currentPath = strings.Join(currentPathItems, "-")
				contents[currentPath] = nil
			}
		} else {
			contents[currentPath] = append(contents[currentPath], line)
		}
	}

	candidate := traverseDirectory("/", contents, sizes, 1)

	const MINIMUM_TO_DELETE = 2536714

	for _, size := range sizes {
		if size > MINIMUM_TO_DELETE {
			if size < candidate {
				candidate = size
			}
		}
	}

	fmt.Println(candidate)
}

func traverseDirectory(name string, contents map[string][]string, sizes map[string]int, level int) int {
	size := 0

	for _, item := range contents[name] {
		tmp := strings.Split(item, " ")

		if tmp[0] == "dir" {
			dirSize := traverseDirectory(name+"-"+tmp[1], contents, sizes, level+1)
			size += dirSize
		} else {
			fileSize, _ := strconv.Atoi(tmp[0])
			size += fileSize
		}
	}

	sizes[name] = size

	return size
}
