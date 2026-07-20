package main

import (
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

	var validPaths []string = []string{}

	// use new modified input to search all valid pipe paths
	fileLength := len(input)
	previousIndex := 0
	// map will store string index where a collision was found and appended string at that point
	collisionMap := make(map[int]string)
    
	// traverse all paths and check if they are valid
	for i := 1; i < fileLength; i++ {
		re := regexp.MustCompile(`\|`)
		pipePositions := re.FindAllStringIndex(input[i-1], -1)

	}
}
