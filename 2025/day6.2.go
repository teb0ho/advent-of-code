package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func trashCompactor2() {
	file, err := os.Open("./files/day6.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := []string{}
	numbers := [][]int{}
	digits := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		list = append(list, line)
	}

	var ops []string

	var nums []string

	for i, val := range list {
		nums = strings.Fields(val)
		ints := []int{}

		if i != len(list)-1 {
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				ints = append(ints, n)

			}

			numbers = append(numbers, ints)
			digits = append(digits, nums)
		} else {
			ops = strings.Fields(list[len(list)-1])
		}
	}
	largestDigits := make([]int, len(digits[0]))

	for i := 0; i < len(digits); i++ {
		for j, num := range digits[i] {
			if len(num) > largestDigits[j] {
				largestDigits[j] = len(num)
			}
		}
	}

	finalSum := len(numbers[0])
	totalArray := make([]int, finalSum)

	for i := 0; i < len(numbers); i++ {
		for j, num := range numbers[i] {
			operator := ops[j]

			switch operator {
			case "+":
				totalArray[j] += num
			case "*":
				if totalArray[j] == 0 {
					totalArray[j] = num
				} else {
					totalArray[j] *= num
				}
			}

		}
	}
	total := 0
	for _, num := range totalArray {
		total += num
	}

	fmt.Println("Total:", total)
}
