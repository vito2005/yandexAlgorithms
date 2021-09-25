package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var max int
	var word string
	dict := make(map[string]int)
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/c/input.txt")
	str := string(file)
	words := strings.Fields(str)
	for _, w := range words {
		_, ok := dict[w]
		if !ok {
			dict[w] = 0
		}
		dict[w]++
		if dict[w] > max {
			max = dict[w]
			word = w
		} else if dict[w] == max && w < word {
			word = w
		}
	}

	fmt.Print(word)
}
