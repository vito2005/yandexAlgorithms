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
	file, err := os.Open("1/g/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	params := []string{}

	for scanner.Scan() {
		params = strings.Fields(scanner.Text())
	}
	N, _ := strconv.Atoi(params[0])
	K, _ := strconv.Atoi(params[1])
	M, _ := strconv.Atoi(params[2])
	q := 0

	q = produceM(N, K, M, q)

	fmt.Println(q)

}

func produceM(N int, K int, M int, q int) int {
	var reminderN, qK, remainderK, qM int

	reminderN = N % K
	qK = N / K

	remainderK = (K % M) * qK
	qM = (K / M) * qK
	if qM == 0 {
		return 0
	}
	q += qM

	if remainderK > 0 && (reminderN+remainderK) >= K {
		return produceM(remainderK+reminderN, K, M, q)
	}

	return q
}
