package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dict := make(map[int]int)

	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/d/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	keysCount, _ := strconv.Atoi(params[0])
	keysMax := strings.Fields(params[1])

	keysPressedCount, _ := strconv.Atoi(params[2])
	keysPressed := strings.Fields(params[3])

	for i := 0; i < keysPressedCount; i++ {
		key, _ := strconv.Atoi(keysPressed[i])
		_, ok := dict[key]
		if !ok {
			dict[key] = 0
		}
		dict[key]++
	}

	for i := 0; i < keysCount; i++ {
		max, _ := strconv.Atoi(keysMax[i])

		if max >= dict[i+1] {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}

}
