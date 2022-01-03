package main

import (
	"fmt"
	"math"
)

//хотим возвратить из функции сумму двух чисел
func main() {
	var a = add(4, 5)  // 9
	var b = add(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(math.Sqrt(float64(add(6, 7))))
	fmt.Println(int64(math.Sqrt(float64(add(6, 7)))))
	fmt.Println(math.Pi * float64(add(8, 9)))
	fmt.Println(int64(math.Pi * float64(add(8, 9))))

	fmt.Println(add2(7, 8))
	fmt.Println(add1(5, 6))
}
func add(x, y int) int {
	return x + y
}

//Фактически мы также могли бы написать, если результат не был именован:
func add2(x, y int) int {
	var z int = x + y
	return z
}

//Возвращаемый результат может быть именован:
func add1(x, y int) (z int) {
	z = x + y
	return // тогда после return можем не писать значение или переменную
}
