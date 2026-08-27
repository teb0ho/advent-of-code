package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./files/day7.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := []string{}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		list = append(list, line)
	}

	laboratoriesPart2(list)
}
