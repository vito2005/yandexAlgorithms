package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1. Special cases/b/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sides := []string{}
	for scanner.Scan() {
		sides = append(sides, scanner.Text())
	}

	var res = "NO"
	a, err := strconv.Atoi(sides[0])
	b, err := strconv.Atoi(sides[1])
	c, err := strconv.Atoi(sides[2])

	if a+b > c && a+c > b && b+c > a {
		res = "YES"
	}

	io.WriteString(os.Stdout,
		res)
}
