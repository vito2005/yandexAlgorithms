package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var res [][]int
	var count int
	file, _ := ioutil.ReadFile("3. Sets/a/input.txt")
	str := string(file)

	numbers := strings.Fields(str)
	for i := 0; i < 10; i++ {
		res = append(res, []int{})
	}

	for _, num := range numbers {
		item, _ := strconv.Atoi(num)
		index := collision(item)
		if !has(item, res[index]) {
			res[index] = append(res[index], item)
		}
	}
	for i := 0; i < 10; i++ {
		count += len(res[i])
	}
	fmt.Println(count)
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
