package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./files/day1.txt")
	position := 50
	totalZeros := 0

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		rotation := scanner.Text()[0]
		count, err := strconv.Atoi(scanner.Text()[1:])

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}

		newPosition := applyRotation(position, rotation, count)

		if newPosition == 0 {
			totalZeros++
		}

		position = newPosition
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %v\n", err)
	}

	fmt.Println(totalZeros)
}

func applyRotation(position int, rotation byte, count int) int {
	if rotation == 'L' {
		for range count {
			if position == 0 {
				position = 99
				continue
			}
			position--
		}
	} else {
		for range count {
			if position == 99 {
				position = 0
				continue
			}
			position++
		}
	}
	return position
}
