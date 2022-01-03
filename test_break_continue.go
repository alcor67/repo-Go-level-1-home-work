package main

import "fmt"

func main() {
	var numbers = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	//сумма только положительных чисел
	for _, value := range numbers {
		if value < 0 {
			continue // переходим к следующей итерации
		}
		sum += value
	}
	fmt.Println("Sum:", sum) // Sum: 27

	var numbers1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum1 = 0

	for _, value := range numbers1 {
		if value > 4 {
			break // если число больше 4 выходим из цикла
		}
		sum1 += value
	}
	fmt.Println("Sum:", sum1) // Sum: 10

}
