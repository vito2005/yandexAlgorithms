package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("2. Linear search/g/input.txt")
	str := string(file)
	numbers := strings.Fields(str)

	max1 := Max(numbers[0], numbers[1])
	max2 := Min(numbers[0], numbers[1])

	min2 := Max(numbers[0], numbers[1])
	min1 := Min(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		current, _ := strconv.Atoi(numbers[i])

		if current >= max1 {
			max2 = max1
			max1 = current
		} else if current >= max2 {
			max2 = current
		}

		if current <= min1 {
			min2 = min1
			min1 = current
		} else if current <= min2 {
			min2 = current
		}
	}

	if max1*max2 > min1*min2 {
		fmt.Println(max2, max1)
	} else {
		fmt.Println(min1, min2)

	}

}

func Max(x, y string) int {
	xint, _ := strconv.Atoi(x)
	yint, _ := strconv.Atoi(y)

	if x < y {
		return yint
	}
	return xint
}

func Min(x, y string) int {
	xint, _ := strconv.Atoi(x)
	yint, _ := strconv.Atoi(y)
	if x > y {
		return yint
	}
	return xint
}
