package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("1. Special cases/f/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = strings.Fields(scanner.Text())
	}
	var h, w int

	h1, _ := strconv.Atoi(params[0])
	w1, _ := strconv.Atoi(params[1])
	h2, _ := strconv.Atoi(params[2])
	w2, _ := strconv.Atoi(params[3])

	combinations := [][]int{{(h1 + h2), max(w1, w2)}, {(h1 + w2), max(h2, w1)}, {(w1 + h2), max(h1, w2)}, {(w1 + w2), max(h1, h2)}}

	min := combinations[0][0] * combinations[0][1]

	for _, value := range combinations {

		if (value[0] * value[1]) <= min {

			min = value[0] * value[1]
			h = value[0]
			w = value[1]
		}
	}

	fmt.Println(h, w)

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
