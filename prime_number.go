package main

import (
	"fmt"
	"math"
)

func main() {

	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	//алгоритм перебора делителей
	//преобразуем int number во float64
	var sqrNumber = math.Sqrt(float64(number))
	//контрольное значение
	var ch = 0
	//Первый делитель
	//i := 2
	for i := 2; i < number; i++ {
		if i <= int(sqrNumber) {
			if number%i == 0 {
				fmt.Println("Число:", number, "составное")
				break
				//fmt.Println(i)
			}
		} else {
			ch = 1
		}
	}
	if ch == 1 {
		fmt.Println("Число:", number, "простое")
	}
}
