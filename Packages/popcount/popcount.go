package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// print the byte array 
	for i, _ := range pc {
		fmt.Println(pc[i]) 
	}
}

func main() {
	var pc [10]byte
	for i := range pc {
		fmt.Println(pc[i])
	}
	
	fmt.Println("Hi");
}
