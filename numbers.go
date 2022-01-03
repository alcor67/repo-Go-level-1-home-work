package main

import (
	"fmt"
)

func main() {

	var number, units, tens, hunrts int
	fmt.Print("Введите 3-х значное число: ")
	fmt.Scan(&number)
	//Определяем число единиц переданного числа
	units = number % 10
	//Определяем число десятков
	tens = ((number - units) / 10) % 10
	//Определяем число сотен
	hunrts = (number - tens*10 - units) / 100
	//Выводим результат
	fmt.Println("    число сотен:", hunrts, "\n",
		"число десятков:", tens, "\n",
		"  число единиц:", units)
}
