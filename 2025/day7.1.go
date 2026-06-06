package main

import (
	"fmt"
	"regexp"
	"strings"
)

func laboratoriesPart1(input []string) {
	position := strings.Index(input[0], "S")
	nextIndices := []int{}

	for i := 1; i < len(input); i++ {
		if nextIndices != nil && len(nextIndices) > 0 {
			var modifiedString string

			for _, nextIndex := range nextIndices {
				modifiedString = input[i]

				if input[i][nextIndex] != '^' {
				}
				modifiedString = modifiedString[:nextIndex] + "|" + modifiedString[nextIndex+1:]
				input[i] = modifiedString
			}

			nextIndices = []int{}
		}

		if i == 1 && position != -1 && !strings.Contains(input[i], "^") {
			input[i] = input[i][:position] + "|" + input[i][position+1:]
		} else if strings.Contains(input[i], "|") {
			re := regexp.MustCompile(`\|`)
			indices := re.FindAllStringIndex(input[i], -1)

			for _, index := range indices {
				nextIndices = append(nextIndices, index[0])
			}

		} else {
			re := regexp.MustCompile(`\^`)
			indices := re.FindAllStringIndex(input[i], -1)

		}
	}

	fmt.Println(input)

}

func seekPipes(input string) {
}
