package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var result []int
	dict := make(map[string]int)
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/b/input.txt")
	str := string(file)
	words := strings.Fields(str)
	for _, w := range words {
		_, ok := dict[w]
		if !ok {
			dict[w] = 0
		} else {
			dict[w]++

		}

		result = append(result, dict[w])
	}
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i], " ")
	}
}
