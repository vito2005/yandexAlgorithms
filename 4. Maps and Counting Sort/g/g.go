package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Costumer struct {
	purchases map[string]int
}

func main() {
	accounts := make(map[string]int)
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/g/input.txt")
	str := string(file)
	records := strings.Split(str, "\n")

	for _, val := range records {
		record := strings.Fields(val)
		if len(record) == 0 {
			break
		}
		operation := record[0]

		switch operation {
		case "BALANCE":
			name := record[1]
			_, ok := accounts[name]
			if !ok {
				fmt.Println("ERROR")
			} else {
				fmt.Println(accounts[name])
			}

		case "DEPOSIT":
			name := record[1]
			sum, _ := strconv.Atoi(record[2])

			checkAccount(name, accounts)

			accounts[name] += sum
		case "WITHDRAW":
			name := record[1]
			sum, _ := strconv.Atoi(record[2])

			checkAccount(name, accounts)

			accounts[name] -= sum
		case "TRANSFER":
			name1 := record[1]
			name2 := record[2]

			sum, _ := strconv.Atoi(record[3])

			checkAccount(name1, accounts)
			checkAccount(name2, accounts)

			accounts[name1] -= sum
			accounts[name2] += sum

		case "INCOME":
			percent, _ := strconv.Atoi(record[1])
			for name, account := range accounts {
				if account > 0 {
					accounts[name] = account + (account*percent)/100
				}
			}

		}

	}
}

func checkAccount(name string, accounts map[string]int) {
	_, ok := accounts[name]
	if !ok {
		accounts[name] = 0
	}

}
