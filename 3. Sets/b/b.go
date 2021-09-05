package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var set [][]int
	var result []int
	file, _ := ioutil.ReadFile("3. Sets/b/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	numbers1 := strings.Fields(params[0])
	numbers2 := strings.Fields(params[1])

	for i := 0; i < 10; i++ {
		set = append(set, []int{})
	}

	for _, num := range numbers1 {
		item, _ := strconv.Atoi(num)
		index := collision(item)
		if !has(item, set[index]) {
			set[index] = append(set[index], item)
		}
	}

	for _, num := range numbers2 {
		item, _ := strconv.Atoi(num)
		index := collision(item)
		if has(item, set[index]) {

			result = append(result, item)

		}
	}

	sort.Ints(result)

	for _, item := range result {
		fmt.Print(item, " ")
	}
}

func has(item int, arr []int) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}

func collision(item int) int {
	return item % 10
}
