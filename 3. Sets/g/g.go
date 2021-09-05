package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var set [][]int
var coll = 100

func main() {
	file, _ := ioutil.ReadFile("3. Sets/g/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	turtleSum, _ := strconv.Atoi(params[0])

	for i := 0; i < coll; i++ {
		set = append(set, []int{})
	}
	for i := 1; i <= turtleSum; i++ {
		ab := strings.Fields(params[i])
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])

		if a+b == turtleSum-1 && a >= 0 && b >= 0 {
			add(a)
		}
	}
	var result int
	for _, s := range set {
		result = result + len(s)
	}
	fmt.Println(result)

}

func has(item int) bool {
	index := collision(item)

	for _, val := range set[index] {
		if val == item {
			return true
		}
	}
	return false
}

func add(item int) {
	index := collision(item)

	if !has(item) {
		set[index] = append(set[index], item)
	}
}

func delete(item int) {
	index := collision(item)
	for i, val := range set[index] {
		if val == item {
			set[index][i] = set[index][len(set[index])-1]
		}
	}
	set[index] = set[index][:len(set[index])-1]
}

func collision(item int) int {
	return item % coll
}
