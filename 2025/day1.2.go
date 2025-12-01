package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findPassword2() {
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

		newPosition, totalZeros1 := applyRotation2(position, rotation, count, totalZeros)

		position = newPosition
		totalZeros = totalZeros1
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error during scanning: %v\n", err)
	}

	fmt.Println(totalZeros)
}

func applyRotation2(position int, rotation byte, count int, totalZeros int) (int, int) {
	if rotation == 'L' {
		for range count {
			if position == 0 {
				position = 99
				continue
			}
			position--
			if position == 0 {
				totalZeros++
			}
		}
	} else {
		for range count {
			if position == 99 {
				position = 0
				totalZeros++
				continue
			}
			position++
		}
	}
	return position, totalZeros
}
