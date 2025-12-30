package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cafeteria() {

	file, err := os.Open("./files/day5.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	secondBatch := false

	firstList := []string{}
	secondList := []string{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			secondBatch = true
			continue
		}

		if !secondBatch {
			firstList = append(firstList, line)
		} else {
			secondList = append(secondList, line)
		}
	}

	finalList := [][]int{}
	for _, num := range secondList {
		numToCheck, _ := strconv.Atoi(num)

		for _, val := range firstList {
			parts := strings.Split(val, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			finalList = append(finalList, []int{start, end})
			// if numToCheck >= start && numToCheck <= end {
			// 	count++
			// 	break
			// }
		}
	}

	fmt.Println(count)
}
