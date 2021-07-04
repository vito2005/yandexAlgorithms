package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1/d/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []int{}
	var result string

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		params = append(params, value)
	}

	a := params[0]
	b := params[1]
	c := params[2]

	var x int
	if a != 0 {
		x = (c*c - b) / a
	}

	if a == 0 && c*c == b {
		result = "MANY SOLUTIONS"
	} else if c >= 0 && a*x+b == c*c && a != 0 {
		result = strconv.Itoa(x)
	} else {
		result = "NO SOLUTION"
	}
	fmt.Println(result)

}
