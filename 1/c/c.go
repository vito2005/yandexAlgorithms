package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("1/c/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	phones := []string{}
	result := []string{}

	for scanner.Scan() {
		phones = append(phones, scanner.Text())
	}

	for i, phone := range phones {
		phones[i] = strings.Join(parsePhone(phone), "")
		if len(phones[i]) == 7 {
			phones[i] = "+7495" + phones[i]
		}
		if i > 0 && phones[i] == phones[0] {
			result = append(result, "YES")
		} else if i > 0 {
			result = append(result, "NO")
		}
	}

	for _, resItem := range result {
		fmt.Println(resItem)
	}

}

func inslice(n string, h []string) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}

func parsePhone(phoneNumber string) []string {
	unnessesarySymbs := []string{"(", ")", "-"}
	parsedPhone := []string{}
	simbs := strings.Split(phoneNumber, "")
	for _, simb := range simbs {
		if !inslice(simb, unnessesarySymbs) {
			if simb == "8" {
				parsedPhone = append(parsedPhone, "+7")
			} else {
				parsedPhone = append(parsedPhone, simb)
			}
		}
	}
	return parsedPhone
}
