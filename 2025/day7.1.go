package main

import (
	"fmt"
	"strings"
)

func laboratoriesPart1(input []string) {
	position := strings.Index(input[0], "S")

	for i := 1; i < len(input); i++ {
		if position != -1 && !strings.Contains(input[i], "^") {
			input[i] = input[i][:position] + "|" + input[i][position+1:]
		}
	}

	fmt.Println(input)

}
