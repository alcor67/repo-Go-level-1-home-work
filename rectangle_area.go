package main

import (
	"fmt"
)

func main() {
	var a, b float32
	fmt.Print("Введите первую сторону прямоугольника: ")
	fmt.Scan(&a)
	//fmt.Print("Первая сторона прямоугольника: \n", a, "\n")
	fmt.Print("Введите вторую сторону прямоугольника: ")
	fmt.Scan(&b)
	//fmt.Print("Вторая сторона прямоугольника: \n", b, "\n")
	fmt.Print("Площадь прямоугольника: ", a*b, "\n")
}
