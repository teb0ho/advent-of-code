package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cafeteria2v2() {
	file, err := os.Open("./files/day5.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := []int{}
	list := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		list = append(list, line)
	}

	var finalList [][2]int
	for _, val := range list {

	}

	// sort results
}
