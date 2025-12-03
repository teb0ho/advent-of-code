package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lobby() {
	file, err := os.Open("./files/day3.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close() // Ensure the file is closed
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		bank := scanner.Text()
		index := 0
		newString := ""
		max := 0
		total := 0
		maxCombined := 0
		highest := ""

		for i, character := range bank {
			number := int(character - '0')

			if i != (len(bank)-1) && number > max {
				max = number
				index = i
				highest += strconv.Itoa(max)
			}
		}

		for i, character := range bank {

			if i != index {
				newString += string(character)
			}
		}

		for character := range newString {
			number := int(character - '0')

			highest += strconv.Itoa(number)
			highestConv, _ := strconv.Atoi(highest)

			if highestConv > maxCombined {
				maxCombined = highestConv
			}
		}

	}

	fmt.Println(total)
}
