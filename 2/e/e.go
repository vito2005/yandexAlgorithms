package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	file, _ := ioutil.ReadFile("2/e/input.txt")
	str := string(file)

	params := strings.Split(str, "\n")

	length, _ := strconv.Atoi(params[0])

	numbers := strings.Fields(params[1])

	var placeIndex, place int
	var max int = 0
	var maxIndex int

	for i := 0; i < length; i++ {
		current, _ := strconv.Atoi(numbers[i])
		if current > max {
			max = current
			maxIndex = i
		}
	}

	for i := maxIndex + 1; i < length-1; i++ {
		current, _ := strconv.Atoi(numbers[i])
		next, _ := strconv.Atoi(numbers[i+1])
		if numbers[i][utf8.RuneCountInString(numbers[i])-1] == 53 && next < current && current <= max {
			var vasRes1 int
			if placeIndex > 0 {
				vasRes, _ := strconv.Atoi(numbers[placeIndex])
				vasRes1 = vasRes
			}

			if vasRes1 < current {
				placeIndex = i
			}
		}
	}

	if placeIndex > 0 {
		vasRes, _ := strconv.Atoi(numbers[placeIndex])
		place = 1
		for i := 0; i < length; i++ {
			current, _ := strconv.Atoi(numbers[i])
			if vasRes < current && i != placeIndex {
				place++
			}
		}
	}
	fmt.Println(place)
}
