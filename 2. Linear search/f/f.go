package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var count int

func main() {
	file, _ := ioutil.ReadFile("2. Linear search/f/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	numbers := strings.Fields(params[1])

	for !isSymmetric(numbers) {
		count++

	}

	if count == 0 {
		fmt.Println(count)
	} else {
		fmt.Println(count)
		for i := count - 1; i >= 0; i-- {
			fmt.Print(numbers[i], " ")
		}
	}
}

func isSymmetric(arr []string) bool {
	length := len(arr)

	for i := 0; i < length; i++ {
		if arr[i] != arr[length-i-1] {
			count++
			return isSymmetric(arr[1:length])
		}
	}
	return true
}
