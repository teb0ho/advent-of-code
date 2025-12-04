package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lobby2() {
	file, err := os.Open("./files/day3.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	// 999999999987
	defer file.Close()

	total := 0
	max := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		stringSample := ""

		for j, _ := range bank {
			for k, _ := range bank {
				for h, _ := range bank {
					for i, character := range bank {
						if i != j && i != k && i != h && k != i && k != j && k != h && h != i && h != j && h != k {
							number := int(character - '0')
							stringSample += strconv.Itoa(number)

							if len(stringSample) == 12 {
								numberConv, _ := strconv.Atoi(stringSample)
								if numberConv > max {
									max = numberConv
								}
								stringSample = ""
							}
						}
					}
				}
			}
		}
		total += max
		fmt.Println(max)
		max = 0

	}
	fmt.Println(total)
}
