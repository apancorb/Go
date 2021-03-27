package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	fmt.Print("here\n")
	return x + x
}

func double1(x int) (result int) {
	defer func(x int) func(y int) {
		fmt.Printf("double(%d) = %d\n", x, result)
		return func(y int) { fmt.Println(y) }
	}(12)(12)
	fmt.Print("here\n")
	return x + x
}

func main() {
	defer func() {
		fmt.Println("End")
	}()
	bigSlowOperation()
	double1(5)
}
