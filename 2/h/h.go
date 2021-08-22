package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var res []int
	file, _ := ioutil.ReadFile("2/h/input.txt")
	str := string(file)
	numbers := strings.Fields(str)

	for _, el := range numbers {
		elint, _ := strconv.Atoi(el)
		res = append(res, elint)
	}
	l := len(res)

	k_rearrange(res, l-1) // чтобы найти максимальное

	k_rearrange(res, l-3) // оставшиеся два максимума

	k_rearrange(res, 2) // два самых маленьких

	if res[l-2]*res[l-3]*res[l-1] > res[0]*res[1]*res[l-1] {
		fmt.Println(res[l-1], res[l-2], res[l-3])
	} else {
		fmt.Println(res[0], res[1], res[l-1])

	}
}

func k_rearrange(seq []int, k int) {
	left := 0
	right := len(seq) - 1

	for left < right {
		x := seq[(right+left)/2]
		eqxfirst := left
		grxfirst := left
		for i := left; i < right+1; i++ {

			now := seq[i]
			if now == x {
				seq[i] = seq[grxfirst]
				seq[grxfirst] = now
				grxfirst++
			} else if now < x {
				seq[i] = seq[grxfirst]
				seq[grxfirst] = seq[eqxfirst]
				seq[eqxfirst] = now
				grxfirst++
				eqxfirst++
			}

		}

		if k < eqxfirst {
			right = eqxfirst - 1
		} else if k >= grxfirst {
			left = grxfirst
		} else {
			return
		}
	}

}
