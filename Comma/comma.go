package main 

import (
	"fmt"
	"strings"
)

// first function comma that uses string concatenation and recursion 
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	var ans string
	var isNegative bool
	// check if it is negative
	if s[0] == '-' {
		isNegative = true
	}

	dot := strings.LastIndex(s, ".") 
	if dot == -1 {
		if isNegative {
			ans = comma(s[1:])
		} else {
			ans = comma(s)
		}
	} else {
		if isNegative {
			ans = "-" + comma(s[1:dot]) + s[dot+1:]
		} else {
			ans = comma(s[:dot]) + s[dot+1:]
		}
	}

	return ans 

//	for i, v := range s {
//		if i != 0 && i % 3 == 0 {
//			buf.WriteString(",")
//		}
//		buf.WriteString(string(v))
//	}
//	return buf.String()
}


func main() {
	var test1 string
	test1 = "abcdefghi"

	test2 := "Hey Boy"

	fmt.Println(comma(test1))
	fmt.Println(comma(test2))

	fmt.Println(comma1(test1))

	fmt.Println(comma1("123"), comma1("123123"), comma1("-123123"), comma1("-123123.0320012"))
}



