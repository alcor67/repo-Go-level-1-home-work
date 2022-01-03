package main

import (
	"fmt"
)

func main() {
	a := true
	fmt.Println(!a)     // false
	fmt.Println(!!a)    // true
	fmt.Println(!!!a)   // true
	fmt.Println(!!!!a)  // true
	fmt.Println(!!!!!a) // true
}
