package main

import (
	"bufio"
	"fmt"
	"os"
)

func printingDepartment() {
	file, err := os.Open("./files/day3.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	//total := 0
	lineCount := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineCount = append(lineCount, line)
	}

	for i := 0; i < len(lineCount[0]); i++ {
	}
}
