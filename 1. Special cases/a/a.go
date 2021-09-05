package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("1. Special cases/a/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}
	var result int

	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	temps := strings.Fields(params[0])

	troom, _ := strconv.Atoi(temps[0])
	tcond, _ := strconv.Atoi(temps[1])
	mode := params[1]

	result = getTemp(troom, tcond, mode)
	fmt.Println(result)
}

func getTemp(troom int, tcond int, mode string) int {
	switch mode {
	case "heat":
		{
			if troom <= tcond {
				return tcond
			} else {
				return troom
			}
		}
	case "freeze":
		{
			if troom <= tcond {
				return troom
			} else {
				return tcond
			}
		}
	case "auto":
		{
			return tcond
		}
	case "fan":
		{
			return troom
		}
	default:
		{
			return troom
		}
	}

}
