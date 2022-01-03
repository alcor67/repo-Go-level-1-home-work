package main

import (
	"fmt"
)

func main() {
	s := "Hello, " + "World!"
	fmt.Println(s)
	fmt.Println(len(s))
	s = "abcd"
	fmt.Println(len(s))                 // Длина 4
	fmt.Println(s[0], s[1], s[2], s[3]) // 97 98 99 100
	s = "🤓"
	fmt.Println(len(s))                 // Длина тоже 4
	fmt.Println(s[0], s[1], s[2], s[3]) // 240 159 164 147
}
