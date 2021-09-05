package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("3. Sets/j/input.txt")
	str := string(file)
	params := strings.Split(str, "\n")

	tdn := strings.Fields(params[0])

	t, _ := strconv.Atoi(tdn[0])
	d, _ := strconv.Atoi(tdn[1])
	n, _ := strconv.Atoi(tdn[2])

	postRect := [4]int{0, 0, 0, 0}
	for i := 1; i <= n; i++ {
		postRect = extend(postRect, t)
		coord := strings.Fields(params[i])
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		navRect := extend([4]int{x + y, x + y, x - y, x - y}, d)
		postRect = intersect(postRect, navRect)
	}

	var points [][]int

	for xPlusY := postRect[0]; xPlusY <= postRect[1]; xPlusY++ {
		for xMinusY := postRect[2]; xMinusY <= postRect[3]; xMinusY++ {
			if (xPlusY+xMinusY)%2 == 0 {
				x := (xPlusY + xMinusY) / 2
				y := xPlusY - x
				points = append(points, []int{x, y})
			}
		}
	}

	fmt.Println(len(points))
	for _, point := range points {
		fmt.Println(point[0], point[1])
	}
}

func min(min1 int, min2 int) int {
	if min1 <= min2 {
		return min1
	}
	return min2
}

func max(max1 int, max2 int) int {
	if max1 >= max2 {
		return max1
	}
	return max2
}

func extend(rect [4]int, t int) [4]int {
	minPlus := rect[0]
	maxPlus := rect[1]
	minMinus := rect[2]
	maxMinus := rect[3]

	return [4]int{minPlus - t, maxPlus + t, minMinus - t, maxMinus + t}
}

func intersect(rect1 [4]int, rect2 [4]int) [4]int {
	return [4]int{max(rect1[0], rect2[0]), min(rect1[1], rect2[1]), max(rect1[2], rect2[2]), min(rect1[3], rect2[3])}
}
