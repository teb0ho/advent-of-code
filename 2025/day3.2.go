package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func lobby2() {
	file, err := os.Open("./files/day3.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		stringSample := ""
		left := 12
		max := 0
		// 14 - 12 = 2

		for i := 0; i < len(bank); i++ {
			if left == 0 {
				break
			}
			character := bank[i]

			number := int(character - '0')

			if len(bank)-i > left {
				if number == 9 {
					stringSample += strconv.Itoa(number)
					left--
					max = 0
				} else if number != 9 {
					moves := len(bank) - i - left + i
					index := 0
					indexList := []int{}

					for j := i; j <= moves; j++ {
						character := bank[j]
						number = int(character - '0')

						if number == max {
							max = number
							indexList = append(indexList, index)
							indexList = append(indexList, j)
							index = j
						} else if number > max {
							max = number
							index = j
						}

					}

					if max != 0 {
						stringSample += strconv.Itoa(max)

						if len(indexList) > 0 {
							sort.Ints(indexList)
							i = indexList[0]
						} else if index != 0 {
							i = index
						}
					} else {
						stringSample += strconv.Itoa(number)
					}
					left--
					max = 0

				} else {
					if max != 0 {
						stringSample += strconv.Itoa(max)
						max = 0
					} else {
						stringSample += strconv.Itoa(number)
					}
					left--
					max = 0
				}
			} else if len(bank)-i == left {
				if len(stringSample) != 12 {
					if max > number {
						stringSample += strconv.Itoa(max)
						max = 0
					} else {
						stringSample += strconv.Itoa(number)

						max = 0
					}
					left--
				}
			} else {
				if number > max {
					max = number
				}
			}

			if len(stringSample) == 12 {
				fmt.Println(stringSample)
				numberConv, _ := strconv.Atoi(stringSample)
				total += numberConv
				stringSample = ""
				max = 0
			}
		}
	}
	fmt.Println("")
	fmt.Println(total)
}
