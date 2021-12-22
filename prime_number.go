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
	//fmt.Println(getPrime(number))
	if number < 2 {
	} else {
		fmt.Println(2)
		for j := 2; j <= number; j++ {
			if getPrime(j) {
				fmt.Println(j)
			}
		}
	}
}

func getPrime(number int) bool {
	var sqrNumber = math.Sqrt(float64(number))
	//контрольное значение
	var ch = false
	//Первый делитель
	//i := 2
	for i := 2; i < number; i++ {
		if i <= int(sqrNumber) {
			if number%i == 0 {
				break
			}
		} else {
			ch = true
		}
	}
	if ch {
		//fmt.Println("Число:", number, "простое")
		return true
	}
	return false
}
