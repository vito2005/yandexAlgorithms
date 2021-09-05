package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"regexp"
	"strings"
)

var set [][]uint64
var coll uint64 = 100000

func main() {
	file, _ := ioutil.ReadFile("3. Sets/d/input.txt")
	str := strings.TrimSpace(string(file))
	re := regexp.MustCompile(`[\s]+`)
	split := re.Split(str, -1)
	var i uint64 = 0
	for i < coll {
		i++
		set = append(set, []uint64{})
	}

	for _, w := range split {
		if w != "" {
			wordInt := hash(w)
			add(wordInt)
		}
	}

	var result int
	for _, setItem := range set {
		result = result + len(setItem)
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
