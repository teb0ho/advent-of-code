package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type PathIndex struct {
	concatenatedString string
	indexX             int
	indexY             int
}

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
	// traverse all paths and check if they are valid
	// LLDDRR
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
					if string(input[row][col]) == "|" {
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

	dfs(tree, "1,7", make(map[string]bool), &count)

	// for _, path := range paths {
	// 	if len(path.concatenatedString) == fileLength {
	// 		count++
	// 	}
	// }

	fmt.Println(count)
}

func dfs(tree map[string][]string, start string, visited map[string]bool, count *int) {

	if visited[start] {
		return
	}
	visited[start] = true

	if strings.Contains(start, "15,") {
		*count++
	}

	for _, neighbor := range tree[start] {
		dfs(tree, neighbor, visited, count)
	}
}

func checkCollictionOrTraverseVertically(input []string, i int, indexX int, paths []PathIndex, parentIndex int) []PathIndex {
	indexOfReferenceElement := findElementIndex(paths, indexX, parentIndex)

	if indexOfReferenceElement == -1 {
		return paths
	}

	if string(input[i][indexX-1]) == "|" {
		paths = append(paths, PathIndex{concatenatedString: paths[indexOfReferenceElement].concatenatedString + string(input[i][indexX-1]), indexX: indexX - 1, indexY: i})
	}

	if string(input[i][indexX+1]) == "|" {
		paths = append(paths, PathIndex{concatenatedString: paths[indexOfReferenceElement].concatenatedString + string(input[i][indexX+1]), indexX: indexX + 1, indexY: i})
	}

	// remove the parent path from the paths slice

	paths = slices.Delete(paths, indexOfReferenceElement, indexOfReferenceElement+1)

	return paths
}

func findElementIndex(paths []PathIndex, indexX int, parentIndex int) int {
	indexOfReferenceElement := -1

	for i, path := range paths {
		if path.indexX == indexX && path.indexY == parentIndex {
			indexOfReferenceElement = i
			break
		}
	}

	return indexOfReferenceElement
}
