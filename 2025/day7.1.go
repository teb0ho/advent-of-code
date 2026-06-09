package main

import (
	"fmt"
	"regexp"
	"strings"
)

func laboratoriesPart1(input []string) {
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
		}
	}

	fmt.Println(input)

}

func seekPipes(input string) {}
