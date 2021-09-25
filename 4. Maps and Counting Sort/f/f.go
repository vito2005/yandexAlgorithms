package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Costumer struct {
	purchases map[string]int
}

func main() {
	dict := make(map[string]map[string]int)
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/f/input.txt")
	str := string(file)
	records := strings.Split(str, "\n")
	var keys []string

	for _, val := range records {
		record := strings.Fields(val)
		if len(record) == 0 {
			break
		}
		name := record[0]
		purchaseName := record[1]
		count, _ := strconv.Atoi(record[2])

		_, ok := dict[name]
		if !ok {
			dict[name] = make(map[string]int)
		}

		currentName := dict[name]

		_, ok2 := currentName[purchaseName]
		if !ok2 {
			currentName[purchaseName] = 0
		}
		currentName[purchaseName] += count
	}

	for k := range dict {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k + ":")
		var ps []string

		for p := range dict[k] {
			ps = append(ps, p)
		}

		sort.Strings(ps)

		for _, p := range ps {
			fmt.Println(p, dict[k][p])
		}
	}
}
