package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"strings"
)

var set [][]uint64
var coll uint64 = 100000

func main() {
	file, _ := ioutil.ReadFile("3. Sets/f/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	var i uint64 = 0
	for i < coll {
		i++
		set = append(set, []uint64{})
	}
	second := strings.Split(params[1], "")
	for i := 0; i < len(second)-1; i++ {
		genItem := second[i] + second[i+1]
		wordInt := hash(genItem)
		add(wordInt)
	}

	var result int

	first := strings.Split(params[0], "")
	for i := 0; i < len(first)-1; i++ {
		genItem := first[i] + first[i+1]
		wordInt := hash(genItem)
		if has(wordInt) {
			result++
		}
	}

	fmt.Println(result)

}

func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

func has(item uint64) bool {
	index := collision(item)

	for _, val := range set[index] {
		if val == item {
			return true
		}
	}
	return false
}

func add(item uint64) {
	index := collision(item)

	if !has(item) {
		set[index] = append(set[index], item)
	}
}

func collision(item uint64) uint64 {
	return item % coll
}
