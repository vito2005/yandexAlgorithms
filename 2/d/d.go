package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2/d/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []string{}
	var result int

	for scanner.Scan() {
		numbers = strings.Fields(scanner.Text())
	}

	for i := 1; i < len(numbers)-1; i++ {
		current, _ := strconv.ParseFloat(numbers[i], 64)
		prev, _ := strconv.ParseFloat(numbers[i-1], 64)
		next, _ := strconv.ParseFloat(numbers[i+1], 64)
		if current > prev && current > next {
			result++
		}
	}

	fmt.Println(result)
}
