package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

type DictValues = []string

func main() {
	mistakes := 0
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/i/input.txt")
	str := string(file)
	records := strings.Split(str, "\n")
	n, _ := strconv.Atoi(records[0])
	dict := make(map[string][]string)

	for i := 1; i <= n; i++ {
		sLower := strings.ToLower(records[i])
		_, ok := dict[sLower]
		if !ok {
			dict[sLower] = []string{}
		}

		dict[sLower] = append(dict[sLower], records[i])
	}
	text := strings.Fields(records[n+1])

	for _, word := range text {

		wordLower := strings.ToLower(word)

		if _, ok := dict[wordLower]; ok {
			if !has(word, dict[wordLower]) {
				mistakes++
			}
		} else {
			stresses := 0
			for _, c := range word {
				if unicode.IsUpper(c) {
					stresses++
				}
			}
			if stresses != 1 {
				mistakes++

			}
		}
	}
	fmt.Println(mistakes)
}

func has(item string, arr []string) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}
