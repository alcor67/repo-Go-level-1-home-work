package main

import (
	"fmt"

	//https://asaskevich.github.io/2015/02/09/validatsiia-dannykh-v-go-pri-pomoshchi-govalidator/
	//https://sprosi.pro/questions/903225/proverte-yavlyaetsya-li-stroka-int-golang
	//https://www.8host.com/blog/preobrazovanie-tipov-dannyx-v-go/
	//https://golang-blog.blogspot.com/2020/04/int-int64-string-convertions.html

	"strconv" //https://golang-blog.blogspot.com/2020/04/float-string-covert.html
	//https://golang-blog.blogspot.com/2020/04/float-string-covert.html
	//https://www.8host.com/blog/preobrazovanie-tipov-dannyx-v-go/
	//https://golang-blog.blogspot.com/2020/04/int-int64-string-convertions.html
)

func messageError() {
	fmt.Println("Ввод данных не корректен:") //сообщение об ошибке
}

func valid1(numberString string) bool {
	//конвертация строки в число метод strconv.ParseFloat(string, 64)
	number, err := strconv.ParseFloat(numberString, 64)
	if err != nil {
		_ = number
		messageError() //сообщение об ошибке
		//сообщение о типе ошибки
		fmt.Println("type of error: value is not a number")
		//системное сообщение об ошибке
		fmt.Println(err.Error())
		return false
	} else {
		//fmt.Println("число valid")
		return true
	}
}

func main() {

	// ввод двух значений и оператора
	fmt.Print("введите через пробел 2 значения и оператор\n",
		"{значение_1 оператор значение_2}\n")
	var numberString1, numberString2 string
	var operator string
	var number1, number2 float64
	fmt.Scanln(&numberString1, &operator, &numberString2)
	//валидация первого числа
	if !valid1(numberString1) {
		return
	}
	//конвертация строки в число метод strconv.ParseFloat(string, 64)
	number1, err := strconv.ParseFloat(numberString1, 64)
	_ = err
	//валидация второго числа
	if !valid1(numberString2) {
		return
	}
	//конвертация строки в число метод strconv.ParseFloat(string, 64)
	number2, err1 := strconv.ParseFloat(numberString2, 64)
	_ = err1
	fmt.Println("значение 1: ", number1)
	fmt.Println("  оператор: ", operator)
	fmt.Println("значение 2: ", number2)
	switch operator {
	case "+":
		//применяю float32 чтобы избавиться от цифрового хвоста в значении результата
		//кстати, сам догадался
		fmt.Println(" результат: ", float32(number1+number2))
	case "-":
		fmt.Println(" результат: ", float32(number1-number2))
	case "*":
		fmt.Println(" результат: ", float32(number1*number2))
	case "/":
		//todo: проверка деления на ноль
		if number2 == 0 {
			messageError() //сообщение об ошибке
			//сообщение о типе ошибки
			fmt.Println("type of error: division by zero")
		} else {
			fmt.Println(" результат: ", float32(number1/number2))
		}
	case "%":
		fmt.Println("операция взятия остатка от деления")
		fmt.Print("данные и результат будут приведены к целочисленному типу\n",
			"путем отбрасывания дробной части\n")
		if number2 < 1 {
			messageError() //сообщение об ошибке
			//сообщение о типе ошибки
			fmt.Println("type of error: integer divide by zero")
		} else {
			fmt.Println(" результат: ", int(int(number1)%int(number2)))
		}
	default:
		messageError() //сообщение об ошибке
		//сообщение о типе ошибки
		fmt.Println("type of error: unknown operation")
	}
}
