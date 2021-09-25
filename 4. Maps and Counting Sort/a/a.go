package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dict := make(map[string]string)
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/a/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	dictAmount, _ := strconv.Atoi(params[0])

	for i := 1; i <= dictAmount; i++ {
		currentStr := strings.Fields(params[i])
		dict[currentStr[0]] = currentStr[1]
		dict[currentStr[1]] = currentStr[0]
	}
	fmt.Println(dict[params[dictAmount+1]])

}
