package main

import (
	"fmt"
	"strings"
)

func laboratoriesPart1(input []string) {
	sPosition := strings.Index(input[0], "S")

	for i := 1; i < len(input); i++ {
		if string(input[i][sPosition]) != "^" {
			input[i] = input[i][:sPosition] + "|" + input[i][sPosition+1:]
		}
	}

	fmt.Println(input)

}

func seekPipes(input string) {}
