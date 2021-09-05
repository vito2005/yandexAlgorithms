package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var arr []int
	file, _ := ioutil.ReadFile("3. Sets/e/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	numbers1 := strings.Fields(params[0])
	numbers2 := []byte(params[1])

	for _, num := range numbers2 {
		n := string(num)

		hasInDigits := false

		for _, digit := range numbers1 {
			if n == digit {
				hasInDigits = true
			}
		}

		if !hasInDigits {
			number, _ := strconv.Atoi(n)
			if !contains(arr, number) {
				arr = append(arr, number)
			}

		}
	}
	fmt.Println(len(arr))

}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
