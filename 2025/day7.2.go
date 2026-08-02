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

	// var validPaths []string = []string{}

	// use new modified input to search all valid pipe paths
	fileLength := len(input)
	// previousIndex := 0
	// map will store string index where a collision was found and appended string at that point
	collisionMap := make(map[[2]int]string)
	outputString := ""
	index := 0
	// traverse all paths and check if they are valid
	for i := 1; i < fileLength; i++ {
		if i == 1 {
			re := regexp.MustCompile(`\|`)
			pipePositions := re.FindStringIndex(input[i])
			index = pipePositions[0]
			outputString += string(input[i][index])
		} else {
			if string(input[i][index]) == "|" {
				outputString += string(input[i][index])
			} else if len(collisionMap) == 0 && string(input[i][index]) == "^" {
				checkCollictionOrTraverseVertically(input, i, index, collisionMap, outputString)
			} else {
				newMap := make(map[[2]int]string)
				for collisionIndex, collisionString := range collisionMap {
					checkCollictionOrTraverseVertically(input, i, collisionIndex[1], newMap, collisionString)
				}
				collisionMap = newMap
			}
		}

	}
	fmt.Println("Collision Map: ", collisionMap)
}

func checkCollictionOrTraverseVertically(input []string, i int, index int, collisionMap map[[2]int]string, outputString string) {
	str := ""
	if string(input[i][index]) == "^" {
		if string(input[i][index-1]) == "|" {
			collisionMap[[2]int{i, index - 1}] = outputString + string(input[i][index-1])
		}

		if string(input[i][index+1]) == "|" {
			collisionMap[[2]int{i, index + 1}] = outputString + string(input[i][index+1])
		}
	} else if string(input[i][index]) == "|" {
		collisionMap[[2]int{i, index}] = outputString + string(input[i][index])
	}
}
