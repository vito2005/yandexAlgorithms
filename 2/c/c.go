package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2/c/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	len, _ := strconv.Atoi(params[0])
	numbers := strings.Fields(params[1])
	x, _ := strconv.Atoi(params[2])
	var resIndex int
	var current, closer float64
	for i := 0; i < len; i++ {
		number, _ := strconv.Atoi(numbers[i])

		if number*x >= 0 {
			current = math.Abs(float64(number - x))
		} else {
			current = math.Abs(float64(number)) + math.Abs(float64(x))
		}

		if i == 0 {
			closer = current
		}

		if current <= closer {
			closer = current
			resIndex = i
		}

	}

	fmt.Println(numbers[resIndex])
}
