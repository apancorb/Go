// Exercise 1.1: Modify the echo program to als o pr int os.Args[0], the name of the command
// that invoked it
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
