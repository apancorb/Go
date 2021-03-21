package main

import (
	"fmt"
)

func pointerarr(arr *[3]int) {
	arr[0] = 1111111
}

func findDuplicates(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i != j && s[i] == s[j] {
				copy(s[j:], s[j+1:])
				fmt.Println(s, i, j)
			}
		}

	}
	return s
}

func main() {
	var q [3]int = [...]int{1, 2, 3}
	s := []int{2, 3, 4, 2, 2, 3, 33, 3}
	f := [2]int{11, 22}
	pointerarr(&q)
	fmt.Println(q)
	fmt.Println(f)
	s = findDuplicates(s)
	fmt.Println(s)
}
