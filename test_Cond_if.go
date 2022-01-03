package main

import "fmt"

func main() {

	a := 8
	b := 8
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}
}
