package main

import "fmt"

func main() {
	list := make([]int, 4, 4)

	list = Append(list, 5)

	fmt.Printf(
		"slice: %d\nlen: %d\ncap %d\n",
		list,
		len(list),
		cap(list),
	)
}

func Append(list []int, element int) []int { // ...args
	var result []int

	newLen := len(list) + 1

	if newLen <= cap(list) {
		result = list[:newLen]
	} else {
		resultCap := newLen

		if resultCap < 2*len(list) {
			resultCap = 2 * len(list)
		}

		result = make([]int, newLen, resultCap)
		copy(result, list)
	}

	result[len(list)] = element

	return result
}
