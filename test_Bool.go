package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 7
	fmt.Println(a > b, a < b, a == b) // false true false
	if a < b {
		fmt.Println("Hello!")
	}
}
