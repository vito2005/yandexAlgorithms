package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1. Special cases/h/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	n, _ := strconv.Atoi(params[2])
	m, _ := strconv.Atoi(params[3])

	minA := (n-1)*a + n
	maxA := minA + 2*a

	minB := (m-1)*b + m
	maxB := minB + 2*b

	min := maxFunc(minA, minB)

	max := minFunc(maxA, maxB)

	if min <= max {
		fmt.Println(min, max)
	} else {
		fmt.Println("-1")
	}
}

func minFunc(min1 int, min2 int) int {
	if min1 <= min2 {
		return min1
	}
	return min2
}

func maxFunc(max1 int, max2 int) int {
	if max1 >= max2 {
		return max1
	}
	return max2
}
