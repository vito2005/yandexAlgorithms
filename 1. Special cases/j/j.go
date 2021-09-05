package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1. Special cases/j/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}
	a, _ := strconv.ParseFloat(params[0], 64)
	b, _ := strconv.ParseFloat(params[1], 64)
	c, _ := strconv.ParseFloat(params[2], 64)
	d, _ := strconv.ParseFloat(params[3], 64)
	e, _ := strconv.ParseFloat(params[4], 64)
	f, _ := strconv.ParseFloat(params[5], 64)

	var x float64
	var y float64

	if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 && f == 0 {
		fmt.Println("5")
	} else if a*d == b*c && a*f != e*c {
		fmt.Println("0")
	} else if (a == 0 && b == 0 && e != 0) || (c == 0 && d == 0 && f != 0) {
		fmt.Println("0")
	} else if (a == 0 && c == 0 && b*f != d*e) || (b == 0 && d == 0 && a*f != c*e) {
		fmt.Println("0")
	} else if a*d == b*c && a*f == c*e {
		if b == 0 && d == 0 {
			if a != 0 && c != 0 {
				fmt.Println("3", e/a)
			} else if a == 0 {
				if e == 0 {
					fmt.Println("3", f/c)
				}
			} else if c == 0 {
				if f == 0 {
					fmt.Println("3", e/a)
				}
			}
		} else if a == 0 && c == 0 {
			if b != 0 {
				fmt.Println("4", e/b)
			} else if d != 0 {
				fmt.Println("4", f/d)
			}
		} else if b != 0 {
			fmt.Println("1", -a/b, e/b)
		} else if d != 0 {
			fmt.Println("1", -c/d, f/d)
		}
	} else {
		y = (a*f - c*e) / (a*d - c*b)
		x = (e*d - b*f) / (a*d - b*c)
		fmt.Println("2", x, y)
	}
}
