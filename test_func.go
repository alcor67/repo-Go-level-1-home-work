package main

import "fmt"

/*
хотим выполнить в программе нашу функцию hello,
нам надо вызвать ее в функции main:
*/
func main() {
	hello()
	hello()
	hello()
	add(4, 5)  // x + y = 9
	add(20, 6) // x + y = 26
	add1(1, 2, 3.4, 5.6, 1.2)
}

func hello() {
	fmt.Println("Hello World")
}

func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)
}

func add1(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}
