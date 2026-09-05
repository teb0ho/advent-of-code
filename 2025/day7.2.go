package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
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

	var tree = make(map[string][]string)

	var nodes []int
	nodeKeys := []string{}
	for row := 1; row < len(input); row++ {
		if row == 1 {
			re := regexp.MustCompile(`\|`)
			pipePositions := re.FindStringIndex(input[row])
			column := pipePositions[0]
			tree[fmt.Sprintf("%d,%d", row, column)] = []string{}
			nodes = []int{column}
			nodeKeys = []string{fmt.Sprintf("%d,%d", row, nodes[0])}
		} else {
			if len(nodes) != 0 {
				newNodes := []int{}
				newNodeKeys := []string{}
				for _, col := range nodes {
					if string(input[row][col]) == "|" && len(input)-1 == row {
						tree[fmt.Sprintf("%d,%d", row, col)] = []string{}
						continue
					} else if string(input[row][col]) == "|" {
						for _, nk := range nodeKeys {
							part, _ := strconv.Atoi(strings.Split(nk, ",")[1])

							if /*strings.Contains(nk, fmt.Sprintf(",%d", col))*/ slices.Contains(nodes, part) {
								tree[nk] = []string{fmt.Sprintf("%d,%d", row, col)}
							}
						}
						newNodeKeys = append(newNodeKeys, fmt.Sprintf("%d,%d", row, col))

					} else if string(input[row][col]) == "^" {
						for _, nk := range nodeKeys {
							part, _ := strconv.Atoi(strings.Split(nk, ",")[1])

							if /*strings.Contains(nk, fmt.Sprintf(",%d", col))*/ slices.Contains(nodes, part) {
								tree[nk] = []string{fmt.Sprintf("%d,%d", row, col-1), fmt.Sprintf("%d,%d", row, col+1)}
							}
						}
						newNodes = append(newNodes, col-1, col+1)
						newNodeKeys = append(newNodeKeys, fmt.Sprintf("%d,%d", row, col+1), fmt.Sprintf("%d,%d", row, col-1))
					}
				}

				nodes = newNodes
				nodeKeys = newNodeKeys
			}
		}
	}

	count := countPaths(tree, "1,7")

	fmt.Println(count)
}

func countPaths(tree map[string][]string, node string) int {
	if len(tree[node]) == 0 {
		return 1
	}

	count := 0

	for _, child := range tree[node] {
		count += countPaths(tree, child)
	}

	return count
}
