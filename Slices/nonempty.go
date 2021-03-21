package main

import (
	"fmt"
)

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

/* func nonempty1(strings []string) []string {
	out := strings[:0]

}
*/
func main() {
	test1 := []string{"Hey", "", "There"}
	res := nonempty(test1)
	fmt.Printf("Original %q %d %d | nonempty %q %d %d\n", test1, len(test1), cap(test1), res, len(res), cap(res))
	test1[0] = "hhhhhh"
	fmt.Printf("Original %q %d %d | nonempty %q %d %d\n", test1, len(test1), cap(test1), res, len(res), cap(res))

	test2 := []string{"I", "", "have"}
	test2 = nonempty(test2)
	fmt.Printf("Original %q %d %d\n", test2, len(test2), cap(test2))

	// slice tests
	a := []int{1, 2, 3}
	fmt.Printf("slice = %v\n", a)
	var b []int
	b = a
	b[0] = 23
	fmt.Printf("slice = %v and copy = %v\n", a, b)
	b = append(b, 123)
	fmt.Printf("%v %d %d | %v %d %d\n", a, len(a), cap(a), b, len(b), cap(b))
	b[0] = 21213
	fmt.Printf("%v %d %d | %v %d %d\n", a, len(a), cap(a), b, len(b), cap(b))

}
