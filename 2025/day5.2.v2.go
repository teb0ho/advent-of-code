package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
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
		parts := strings.Split(val, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		finalList = append(finalList, [2]int{start, end})
	}

	// sort results

	sort.Slice(finalList, func(i, j int) bool {
		return finalList[i][0] < finalList[j][0]
	})

	// remove duplicates
	for i := 0; i < len(finalList); i++ {
		for j := i + 1; j < len(finalList); j++ {
			if finalList[i][0] == finalList[j][0] && finalList[i][1] == finalList[j][1] {
				finalList = slices.Delete(finalList, j, j+1)
			}
		}
	}

	// fix overlaps
	for i := 0; i < len(finalList); i++ {
		if i+1 < len(finalList) {
			if finalList[i][1] > finalList[i+1][0] && finalList[i][1] < finalList[i+1][1] {
				startTemp := finalList[i+1][0]
				finalList[i][1] = startTemp
				finalList[i+1][0] = startTemp + 1
			}
			if finalList[i][1] == finalList[i+1][0] && finalList[i][1] < finalList[i+1][1] {
				finalList[i+1][0] = finalList[i+1][0] + 1
			}
		}
	}

	// calculate total
	total := 0
	for _, val := range finalList {
		answer := val[1] - val[0] + 1
		total += answer
	}

	fmt.Printf("Total: %d\n", total)
}

// 391425739187673
// 391425739187688
