package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("2/j/input.txt")

	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	var left float64 = 30
	var right float64 = 4000

	n, _ := strconv.Atoi(params[0])
	prev, _ := strconv.ParseFloat(params[1], 64)

	for i := 2; i <= n; i++ {

		noteData := strings.Fields(params[i])
		note, _ := strconv.ParseFloat(noteData[0], 64)
		place := noteData[1]

		if module(note-prev) < 1/1000000 || note == prev {
			continue
		}

		if place == "closer" {
			if note > prev {
				left = Max(left, (note+prev)/2)

			} else {
				right = Min(right, (note+prev)/2)

			}
		} else {
			if note < prev {
				left = Max(left, (note+prev)/2)
			} else {
				right = Min(right, (note+prev)/2)
			}
		}

		prev = note
	}
	fmt.Println(left, right)

}

func module(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
