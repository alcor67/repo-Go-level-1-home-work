package main

import (
	"fmt"

	//https://asaskevich.github.io/2015/02/09/validatsiia-dannykh-v-go-pri-pomoshchi-govalidator/
	//https://sprosi.pro/questions/903225/proverte-yavlyaetsya-li-stroka-int-golang
	//https://www.8host.com/blog/preobrazovanie-tipov-dannyx-v-go/
	//https://golang-blog.blogspot.com/2020/04/int-int64-string-convertions.html

	"strconv" //https://www.8host.com/blog/preobrazovanie-tipov-dannyx-v-go/
	//https://golang-blog.blogspot.com/2020/04/int-int64-string-convertions.html
)

func messageError() {
	fmt.Println("Вы не ввели положительное 3-х значное число:")
}

func valid1(numberString string) bool {
	//конвертация строки в число метод strconv.Atoi()
	number, err := strconv.Atoi(numberString)
	if err != nil {
		fmt.Println(" input is not digit")
		messageError()
		return false
	} else if number < 100 {
		fmt.Println("number < 100")
		messageError()
		return false
	} else if number > 999 {
		fmt.Println("number > 999")
		messageError()
		return false
	} else {
		fmt.Println("100 =< number <= 999")
		return true
	}
}

func getNumber() string {
	var numberString string
	fmt.Println("Введите положительное 3-х значное число: ")
	fmt.Scanln(&numberString)
	return numberString
}

func main() {

	//Ввод положительного 3-х значного числа
	numberString := getNumber()
	//валидация данных
	if valid1(numberString) {
		//конвертация строки в число strconv.Atoi()
		number, err := strconv.Atoi(numberString)
		//Определяем число единиц переданного числа
		units := number % 10
		//Определяем число десятков
		tens := ((number - units) / 10) % 10
		//Определяем число сотен
		hundreds := (number - tens*10 - units) / 100
		//Выводим результат
		fmt.Println("    число сотен:", hundreds, "\n",
			"число десятков:", tens, "\n",
			//Здесь выводим err как заглушку для метода strconv.Atoi,
			//т.к. требуется переменную err где-то применить, иначе не работает
			"  число единиц:", units, "\n", err)
	}
}
