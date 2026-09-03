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

	var paths []PathIndex = []PathIndex{}

	fileLength := len(input) - 1
	var tree map[string][]string = make(map[string][]string)
	// traverse all paths and check if they are valid
	// LLDDRR
	previous := []int{}
	previousKey := ""
	for i := 1; i < len(input); i++ {
		if i == 1 {
			re := regexp.MustCompile(`\|`)
			pipePositions := re.FindStringIndex(input[i])
			index := pipePositions[0]
			tree[fmt.Sprintf("%d,%d", i, index)] = []string{}
			previous = []int{index}
			previousKey = fmt.Sprintf("%d,%d", i-1, previous[0])
		} else {
			if len(previous) == 0 {

				if string(input[i][previous[0]]) == "|" {
					tree[previousKey] = []string{fmt.Sprintf("%d,%d", i, previous[0])}

				} else if string(input[i][previous[0]]) == "^" {
					tree[previousKey] = []string{fmt.Sprintf("%d,%d", i, previous[0]-1), fmt.Sprintf("%d,%d", i, previous[0]+1)}
					previous = []int{previous[0] - 1, previous[0] + 1}
					previousKey = fmt.Sprintf("%d,%d", i-1, previous[0])
				}
			}

		}
	}

	count := 0

	for _, path := range paths {
		if len(path.concatenatedString) == fileLength {
			count++
		}
	}

	fmt.Println(count)
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
