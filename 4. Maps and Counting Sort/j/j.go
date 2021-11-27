package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := ioutil.ReadFile("4. Maps and Counting Sort/j/input.txt")
	str := string(file)
	records := strings.Split(str, "\n")
	params := strings.Split(records[0], " ")
	n, _ := strconv.Atoi(params[0])
	caseSens := params[1] == "yes"
	stDigit := params[2] == "yes"
	keywords := []string{}

	for i := 1; i <= n; i++ {
		keyword := records[i]
		if !caseSens {
			keyword = strings.ToLower(keyword)
		}
		keywords = append(keywords, keyword)
	}
	type CntPos struct {
		Cnt, Pos int
	}
	cntPosIds := make(map[string][]int)
	wordNumber := 0

	for j := n + 1; j < len(records); j++ {
		line := strings.TrimSpace(records[j])

		line = modifyLine(line)

		words := strings.Fields(line)

		for _, word := range words {
			word = strings.TrimSpace(word)

			if !caseSens {
				word = strings.ToLower(word)
			}

			if has(word, keywords) {
				continue
			}

			if isCorrect(word, stDigit) {

				wordNumber += 1
				_, ok := cntPosIds[word]
				if !ok {
					cntPosIds[word] = []int{0, wordNumber}
				}
				cntPosIds[word][0]++
			}
		}
	}

	var bestWord string

	maxCntPos := []int{0, 0}

	for word, _ := range cntPosIds {
		if cntPosIds[word][0] > maxCntPos[0] {
			maxCntPos = cntPosIds[word]
			bestWord = word
		}
		if cntPosIds[word][0] == maxCntPos[0] && cntPosIds[word][1] < maxCntPos[1] {
			maxCntPos = cntPosIds[word]
			bestWord = word
		}

	}

	fmt.Println(bestWord)

}

func modifyLine(line string) string {

	ans := []rune{}
	for _, c := range line {
		if unicode.IsDigit(c) || unicode.IsLetter(c) || c == 95 {
			ans = append(ans, c)
		} else {
			ans = append(ans, ' ')
		}

	}

	return string(ans)
}

func isCorrect(word string, stDigit bool) bool {

	if isInt(word) {
		return false
	}
	if !unicode.IsDigit(rune(word[0])) || stDigit {
		return true
	}
	return false
}

func has(item string, arr []string) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}

func isInt(val string) bool {
	for _, c := range val {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
