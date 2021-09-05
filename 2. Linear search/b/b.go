package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("2. Linear search/b/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []string{}
	var sequenceType string
	endOfSequence := "-2000000000"

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	i := 1
	for numbers[i] != endOfSequence {

		item, _ := strconv.Atoi(numbers[i])
		itemPrev, _ := strconv.Atoi(numbers[i-1])

		if item < itemPrev {
			if sequenceType == "CONSTANT" {
				sequenceType = "WEAKLY DESCENDING"
			} else if sequenceType == "ASCENDING" || sequenceType == "WEAKLY ASCENDING" {
				sequenceType = "RANDOM"
			} else if sequenceType != "WEAKLY DESCENDING" && sequenceType != "RANDOM" {
				sequenceType = "DESCENDING"
			}
		} else if item > itemPrev {
			if sequenceType == "CONSTANT" {
				sequenceType = "WEAKLY ASCENDING"
			} else if sequenceType == "DESCENDING" || sequenceType == "WEAKLY DESCENDING" {
				sequenceType = "RANDOM"
			} else if sequenceType != "WEAKLY ASCENDING" && sequenceType != "RANDOM" {
				sequenceType = "ASCENDING"
			}
		} else if item == itemPrev {
			if sequenceType == "DESCENDING" {
				sequenceType = "WEAKLY DESCENDING"
			} else if sequenceType == "ASCENDING" {
				sequenceType = "WEAKLY ASCENDING"
			} else if sequenceType != "RANDOM" && sequenceType != "WEAKLY ASCENDING" && sequenceType != "WEAKLY DESCENDING" {
				sequenceType = "CONSTANT"
			}
		}

		i++

	}

	fmt.Println(sequenceType)
}
