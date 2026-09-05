package main

import (
	"fmt"
	"strings"
)

func laboratoriesPart2(input []string) {
	if len(input) < 2 {
		fmt.Println(0)
		return
	}

	paths := make([]uint64, len(input[0]))
	paths[strings.Index(input[0], "S")] = 1

	for row := 1; row < len(input); row++ {
		next := make([]uint64, len(paths))
		for column, count := range paths {
			if count == 0 {
				continue
			}

			if input[row][column] == '^' {
				next[column-1] += count
				next[column+1] += count
			} else {
				next[column] += count
			}
		}
		paths = next
	}

	var count uint64
	for _, pathCount := range paths {
		count += pathCount
	}

	fmt.Println(count)
}
