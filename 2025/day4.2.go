package main

import (
	"bufio"
	"fmt"
	"os"
)

func printingDepartment2() {
	file, err := os.Open("./files/day4.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	lines := []string{}
	matched := [][]int{}
	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for {

		for i := 0; i < len(lines); i++ {
			for j, _ := range lines[i] {
				if rune(lines[i][j]) != '@' {
					continue
				}

				count := 0

				topLeft := []int{i - 1, j - 1}
				top := []int{i - 1, j}
				topRight := []int{i - 1, j + 1}
				left := []int{i, j - 1}
				right := []int{i, j + 1}
				bottomLeft := []int{i + 1, j - 1}
				bottom := []int{i + 1, j}
				bottomRight := []int{i + 1, j + 1}

				if topLeft[0] >= 0 && topLeft[0] < len(lines) && topLeft[1] >= 0 && topLeft[1] < len(lines[0]) {
					if rune(lines[topLeft[0]][topLeft[1]]) == '@' {
						count++
					}
				}

				if top[0] >= 0 && top[0] < len(lines) && top[1] >= 0 && top[1] < len(lines[0]) {
					if rune(lines[top[0]][top[1]]) == '@' {
						count++
					}
				}

				if topRight[0] >= 0 && topRight[0] < len(lines) && topRight[1] >= 0 && topRight[1] < len(lines[0]) {
					if rune(lines[topRight[0]][topRight[1]]) == '@' {
						count++
					}
				}

				if left[0] >= 0 && left[0] < len(lines) && left[1] >= 0 && left[1] < len(lines[0]) {
					if rune(lines[left[0]][left[1]]) == '@' {
						count++
					}
				}

				if right[0] >= 0 && right[0] < len(lines) && right[1] >= 0 && right[1] < len(lines[0]) {
					if rune(lines[right[0]][right[1]]) == '@' {
						count++
					}
				}

				if bottomLeft[0] >= 0 && bottomLeft[0] < len(lines) && bottomLeft[1] >= 0 && bottomLeft[1] < len(lines[i]) {
					if rune(lines[bottomLeft[0]][bottomLeft[1]]) == '@' {
						count++
					}
				}

				if bottom[0] >= 0 && bottom[0] < len(lines) && bottom[1] >= 0 && bottom[1] < len(lines[0]) {
					if rune(lines[bottom[0]][bottom[1]]) == '@' {
						count++
					}
				}

				if bottomRight[0] >= 0 && bottomRight[0] < len(lines) && bottomRight[1] >= 0 && bottomRight[1] < len(lines[0]) {
					if rune(lines[bottomRight[0]][bottomRight[1]]) == '@' {
						count++
					}
				}

				if count < 4 {
					total++
					matched = append(matched, []int{i, j})
				}
			}
		}

		// update matched positions
		for _, pos := range matched {

			i := pos[0]
			j := pos[1]

			lineRunes := []rune(lines[i])
			lineRunes[j] = '.'
			lines[i] = string(lineRunes)
		}

		// add exit condition
		if len(matched) == 0 {
			break
		}

		matched = [][]int{}
	}

	fmt.Println(total)
}
