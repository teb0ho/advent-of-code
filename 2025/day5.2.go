package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func cafeteria2() {
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
		parts := strings.Split(val, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		finalList = append(finalList, [2]int{start, end})

		// for i := start; i <= end; i++ {
		// 	if !slices.Contains(results, i) {
		// 		results = append(results, i)
		// 	}
		// }
	}

	slices.SortFunc(finalList, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	for i := 0; i < len(finalList); i++ {
		currentStart := finalList[i][0]
		currentEnd := finalList[i][1]

		for j := i + 1; j < len(finalList); j++ {
			nextStart := finalList[j][0]
			nextEnd := finalList[j][1]
			if nextStart <= currentEnd && nextEnd <= currentEnd {
				var newList [][2]int

				for k, v := range finalList {
					if k != i {
						newList = append(newList, v)
					}
				}

				finalList = newList
			}

		}
	}

	fmt.Println(len(results))
}
