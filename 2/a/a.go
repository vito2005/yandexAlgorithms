package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2/a/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}
	res := "YES"

	for scanner.Scan() {
		params = strings.Fields(scanner.Text())
	}
	for i := 0; i < (len(params) - 1); i++ {
		item, _ := strconv.Atoi(params[i])
		itemNext, _ := strconv.Atoi(params[i+1])

		if item >= itemNext {
			res = "NO"
		}
	}

	fmt.Println(res)
}
