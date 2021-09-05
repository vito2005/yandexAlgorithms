package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var p1, n1 int
	file, err := os.Open("1. Special cases/e/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}
	const maxVariants = 1000000
	hasAns := false

	for scanner.Scan() {
		params = strings.Fields(scanner.Text())
	}

	k1, _ := strconv.Atoi(params[0])
	m, _ := strconv.Atoi(params[1])
	k2, _ := strconv.Atoi(params[2]) // кв
	p2, _ := strconv.Atoi(params[3]) // под
	n2, _ := strconv.Atoi(params[4]) // эт

	for i := 1; i <= maxVariants; i++ {

		entr2, floor2 := getEntrAndFloor(k2, i, m)

		if p2 == entr2 && n2 == floor2 {

			if hasAns {
				p1New, n1New := getEntrAndFloor(k1, i, m)

				if p1New != p1 {
					p1 = 0
				}

				if n1New != n1 {
					n1 = 0
				}

			} else {
				p1, n1 = getEntrAndFloor(k1, i, m)
			}
			hasAns = true
		}
	}

	if hasAns {
		fmt.Println(p1, n1)
	} else {
		fmt.Println(-1, -1)

	}

}

func getEntrAndFloor(flatNum int, flatsOnFloor int, floors int) (int, int) {
	floorsBefore := ((flatNum - 1) / flatsOnFloor)
	floor := floorsBefore % floors
	entrance := floorsBefore / floors

	return entrance + 1, floor + 1
}
