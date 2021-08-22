package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var mines [][]int
	file, _ := os.Open("2/i/input.txt")

	scanner := bufio.NewScanner(file)
	params := []string{}
	for scanner.Scan() {
		params = append(params, scanner.Text())
	}

	nmk := strings.Fields(params[0])
	n, _ := strconv.Atoi(nmk[0])
	m, _ := strconv.Atoi(nmk[1])
	k, _ := strconv.Atoi(nmk[2])

	for t := 1; t <= k; t++ {
		coords := strings.Fields(params[t])
		i, _ := strconv.Atoi(coords[0])
		j, _ := strconv.Atoi(coords[1])
		arr := []int{i, j}
		mines = append(mines, arr)
	}

	fields := makeFields(n, m, mines)

	for a := 1; a <= n; a++ {
		for b := 1; b <= m; b++ {
			if fields[a][b] == -1 {
				fmt.Print("* ")
			} else {
				fmt.Print(fields[a][b], " ")
			}
		}
		fmt.Println("")
	}
}

func makeFields(n int, m int, mines [][]int) [][]int {
	di := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dj := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	var fields [][]int
	for i := 0; i < n+2; i++ {
		var columns []int
		for j := 0; j < m+2; j++ {
			columns = append(columns, 0)
		}
		fields = append(fields, columns)
	}

	for _, mine := range mines {
		for k := 0; k < len(di); k++ {
			fields[mine[0]+di[k]][mine[1]+dj[k]]++
		}
		fields[mine[0]][mine[1]] = -1
	}

	for _, mine := range mines {
		fields[mine[0]][mine[1]] = -1
	}
	return fields

}
