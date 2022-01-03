package main

import "fmt"

/*
можно передавать и значения переменных, результаты операций или других функций
аргументы в функцию всегда передаются по значению
*/
func main() {
	var a = 8
	fmt.Println("a before: ", a)
	increment(a)
	fmt.Println("a after: ", a)
}
func increment(x int) {
	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}
