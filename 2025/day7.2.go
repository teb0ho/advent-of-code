package main

import (
	"fmt"
	"regexp"
	"strings"
)

func laboratoriesPart2(input []string) {
	sPosition := strings.Index(input[0], "S")

	for i := 1; i < len(input); i++ {
		if i == 1 && string(input[i][sPosition]) != "^" {
			input[i] = input[i][:sPosition] + "|" + input[i][sPosition+1:]
		} else {
			// check pipes in previous row and add pipes to the current row
			re := regexp.MustCompile(`\|`)
			pipePositions := re.FindAllStringIndex(input[i-1], -1)
			// if index has . insert pipe else if it has a caret ^ add pipes to the left and right of the caret
			// and count the pipe and caret collisions
			for _, pos := range pipePositions {
				if string(input[i][pos[0]]) == "." {
					input[i] = input[i][:pos[0]] + "|" + input[i][pos[0]+1:]
				} else if string(input[i][pos[0]]) == "^" {
					out := []rune(input[i])
					out[pos[0]-1] = '|'
					out[pos[0]+1] = '|'
					input[i] = string(out)
				}
			}
		}
	}

	var tree map[string][]string = make(map[string][]string)

	nodes := []int{}
	nodeKey := ""
	for row := 1; row < len(input); row++ {
		if row == 1 {
			re := regexp.MustCompile(`\|`)
			pipePositions := re.FindStringIndex(input[row])
			column := pipePositions[0]
			tree[fmt.Sprintf("%d,%d", row, column)] = []string{}
			nodes = []int{column}
			nodeKey = fmt.Sprintf("%d,%d", row, nodes[0])
		} else {
			if len(nodes) != 0 {
				newNodeKeys := []int{}
				for _, col := range nodes {
					if string(input[row][col]) == "|" && len(input)-1 == row {
						tree[fmt.Sprintf("%d,%d", row, col)] = []string{}
						continue
					} else if string(input[row][col]) == "|" {
						tree[nodeKey] = []string{fmt.Sprintf("%d,%d", row, col)}
						nodeKey = fmt.Sprintf("%d,%d", row, col)
						newNodeKeys = append(newNodeKeys, col)

					} else if string(input[row][col]) == "^" {
						tree[nodeKey] = []string{fmt.Sprintf("%d,%d", row, col-1), fmt.Sprintf("%d,%d", row, col+1)}
						newNodeKeys = append(newNodeKeys, col-1, col+1)
						nodeKey = fmt.Sprintf("%d,%d", row-1, col)
					}
				}

				nodes = newNodeKeys
			}
		}
	}

	count := 0
	count = countPaths(tree, "1,7")

	fmt.Println(count)
}

func countPaths(tree map[string][]string, node string) int {
	// Empty array = we've reached an endpoint
	if len(tree[node]) == 0 {
		return 1
	}

	count := 0

	for _, child := range tree[node] {
		count += countPaths(tree, child)
	}

	return count
}
