package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dict := make(map[int]int)
	sum := 0
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/e/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	blocksCount, _ := strconv.Atoi(params[0])

	for i := 1; i <= blocksCount; i++ {
		block := strings.Fields(params[i])
		w, _ := strconv.Atoi(block[0])
		h, _ := strconv.Atoi(block[1])

		_, ok := dict[w]
		if !ok {
			dict[w] = h
		}

		if dict[w] < h {
			dict[w] = h
		}

	}

	for _, val := range dict {
		sum += val
	}

	fmt.Print(sum)
}
