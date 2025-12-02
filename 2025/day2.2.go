package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func giftShop2() {
	file, err := os.Open("./files/day2.txt")
	total := 0

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close() // Ensure the file is closed

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		vals := scanner.Text()
		parts := strings.SplitSeq(vals, ",")

		for item := range parts {
			ranges := strings.Split(item, "-")
			ranges1Start, _ := strconv.Atoi(ranges[0])
			ranges1End, _ := strconv.Atoi(ranges[1])

			for i := ranges1Start; i <= ranges1End; i++ {
				rangeString := strconv.Itoa(i)

				if len(rangeString)%2 != 0 {
					half := len(rangeString) / 2
					if rangeString[:half] == rangeString[half:] {
						total += i
					}
				} else {

				}
			}
		}
		fmt.Println(total)
	}
}
