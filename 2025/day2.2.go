package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntSet map[int]struct{}

// NewIntSet creates a new empty integer set
func NewIntSet() IntSet {
	return make(IntSet)
}

// Add adds an element to the set
func (s IntSet) Add(element int) {
	s[element] = struct{}{}
}

// Contains checks if an element exists in the set
func (s IntSet) Contains(element int) bool {
	_, exists := s[element]
	return exists
}

// Remove removes an element from the set
func (s IntSet) Remove(element int) {
	delete(s, element)
}

// Size returns the number of elements in the set
func (s IntSet) Size() int {
	return len(s)
}

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

				var results []string
				test := ""

				for j := 1; j < len(rangeString); j++ {
					found := false
					if len(rangeString)%j == 0 {
						for k := 0; k < len(rangeString); k += j {
							end := k + j
							test = rangeString[k:end]
							results = append(results, test)
						}
						allMatched := true
						for _, item := range results {
							if test != item {
								allMatched = false
								break
							}
						}

						if allMatched {
							fmt.Println(results)
							fmt.Println(i)
							total += i
							found = true
							break
						}
						if found {
							break
						}
						results = nil
					} else {
						continue
					}
				}

			}
		}
		fmt.Println(total)
	}
}
