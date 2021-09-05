package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1. Special cases/i/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])
	e, _ := strconv.Atoi(params[4])

	a, b = sortBrickSides(a, b)
	b, c = sortBrickSides(b, c)
	a, b = sortBrickSides(a, b)
	d, e = sortBrickSides(d, e)

	if d >= a && e >= b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func sortBrickSides(side1 int, side2 int) (int, int) {
	if side1 < side2 {
		return side1, side2
	}
	return side2, side1
}
