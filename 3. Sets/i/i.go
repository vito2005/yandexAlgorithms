package main

import (
	//"encoding/binary"

	"fmt"
	"hash/fnv"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("3. Sets/i/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")
	var collConst uint64 = 1000

	set1 := createSet(collConst)
	set2 := createSet(collConst)
	setAllLang := createSet(collConst)

	pupil1LangSum, _ := strconv.Atoi(params[1])

	for i := 2; i <= pupil1LangSum+1; i++ {
		h, lang := hash(params[i])
		set1.add(h, lang)
		setAllLang.add(h, lang)
	}

	for i := 1 + pupil1LangSum + 1; i < len(params)-1; {
		pupilLangs, _ := strconv.Atoi(params[i])
		lenght := pupilLangs + i
		for j := i + 1; j <= lenght; j++ {

			h, lang := hash(params[j])
			if set1.has(h, lang) {
				set2.add(h, lang)
			}
			setAllLang.add(h, lang)

			i++
		}
		set1 = set2
		set2 = createSet(collConst)
		i++

	}

	var result []string
	for _, s := range set1.data {
		for _, l := range s {
			result = append(result, l)
		}
	}

	fmt.Println(len(result))
	for _, l := range result {
		fmt.Println(l)
	}

	var resultAll []string
	for _, s := range setAllLang.data {
		for _, l := range s {
			resultAll = append(resultAll, l)
		}
	}

	fmt.Println(len(resultAll))
	for _, l := range resultAll {
		fmt.Println(l)
	}

}

type Set struct {
	coll uint64
	data [][]string
}

func hash(s string) (uint64, string) {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64(), s
}

func createSet(coll uint64) Set {
	var i uint64 = 0
	set := Set{coll, nil}

	for i < set.coll {
		i++
		set.data = append(set.data, []string{})
	}

	return set
}

func (set Set) has(item uint64, lang string) bool {
	index := set.collision(item)

	for _, val := range set.data[index] {
		if val == lang {
			return true
		}
	}
	return false
}

func (set Set) add(item uint64, lang string) {
	index := set.collision(item)

	if !set.has(item, lang) {
		set.data[index] = append(set.data[index], lang)
	}
}

func (set Set) collision(item uint64) uint64 {
	return item % set.coll
}
