package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Printf("Сумма чисел: %f\n", a+b)
}
