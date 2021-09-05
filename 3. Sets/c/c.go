package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var set [][]int
var coll = 1000

func main() {
	var narr, marr, result []int
	file, _ := ioutil.ReadFile("3. Sets/c/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	nm := strings.Fields(params[0])

	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	for i := 0; i < coll; i++ {
		set = append(set, []int{})
	}

	for i := 1; i <= n; i++ {
		item, _ := strconv.Atoi(params[i])
		add(item)
	}

	for j := n + 1; j <= n+m; j++ {

		item, _ := strconv.Atoi(params[j])

		if has(item) {
			result = append(result, item)
			delete(item)
		} else {

			marr = append(marr, item)
		}
	}

	sort.Ints(result)

	fmt.Println(len(result))
	for _, item := range result {
		fmt.Print(item, " ")
	}
	fmt.Println("")

	for i := 1; i <= n; i++ {
		item, _ := strconv.Atoi(params[i])
		if has(item) {
			narr = append(narr, item)

		}
	}

	fmt.Println(len(narr))
	sort.Ints(narr)
	for _, item := range narr {
		fmt.Print(item, " ")
	}

	fmt.Println("")
	fmt.Println(len(marr))
	sort.Ints(marr)

	for _, item := range marr {
		fmt.Print(item, " ")
	}
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
