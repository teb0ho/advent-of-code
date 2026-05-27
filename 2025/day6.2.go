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

	var stringDigits = make([][]string, len(digits))

	// get substrings of digits by size
	for i := 0; i < len(numbers); i++ {
		index := 0
		stringDigitsRow := []string{}
		for j := 0; j < len(numbers[i]); j++ {
			stringDigitsRow = append(stringDigitsRow, strings.ReplaceAll(list[i][index:index+largestDigits[j]], " ", "0"))
			index += largestDigits[j] + 1
		}
		stringDigits[i] = stringDigitsRow
		stringDigitsRow = []string{}
	}

	finalSum := len(numbers[0])
	totalArray := make([]int, finalSum)
	nums = []string{}

	k := 0
	for k < len(stringDigits[0]) {
		completeNumber := ""

		postition := 0
		for postition < len(stringDigits[0][k]) {
			for i := range stringDigits {

				num := int(stringDigits[i][k][postition] - '0')
				completeNumber += strconv.Itoa(num)
			}
			nums = append(nums, completeNumber)
			completeNumber = ""
			postition++
		}
		finalAnswer := 0
		for _, value := range nums {
			numToAdd := strings.ReplaceAll(value, "0", "")
			operator := ops[k]
			num, _ := strconv.Atoi(numToAdd)
			switch operator {
			case "+":
				finalAnswer += num
			case "*":
				if finalAnswer == 0 {
					finalAnswer = num
				} else {
					finalAnswer *= num
				}
			}
		}
		totalArray = append(totalArray, finalAnswer)
		k++
		//nums = append(nums, completeNumber)
	}

	// for j := 0; j < len(stringDigits); j++ {
	// 	operator := ops[i]
	// 	num, _ := strconv.Atoi(stringDigits[j][i])
	// 	switch operator {
	// 	case "+":
	// 		totalArray[i] += num
	// 	case "*":
	// 		if totalArray[i] == 0 {
	// 			totalArray[i] = num
	// 		} else {
	// 			totalArray[i] *= num
	// 		}
	// 	}
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	for j, num := range numbers[i] {
	// 		operator := ops[j]

	// 		switch operator {
	// 		case "+":
	// 			totalArray[j] += num
	// 		case "*":
	// 			if totalArray[j] == 0 {
	// 				totalArray[j] = num
	// 			} else {
	// 				totalArray[j] *= num
	// 			}
	// 		}

	// 	}
	// }
	// total := 0
	// for _, num := range totalArray {
	// 	total += num
	// }

	fmt.Println("Total:")
}
