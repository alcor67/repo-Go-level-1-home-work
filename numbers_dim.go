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

//Функция вывода сообщения о некорректном вводе данных
func messageError() {
	fmt.Println("Вы не ввели положительное 3-х значное число:")
}

//Функция валидации данных
func valid1(numberString string) bool {
	//конвертация строки в число метод strconv.Atoi()
	//обработка кода ошибки err метода strconv.Atoi()
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

//Функция ввода положительного 3-х значного числа
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
		//https://question-it.com/questions/439825/poluchit-simvol-po-indeksu-stroki
		//strArr := []rune(numberString)
		fmt.Println("    число сотен:", string(numberString[0]), "\n",
			"число десятков:", string(numberString[1]), "\n",
			"  число единиц:", string(numberString[2]), "\n")
	}
}
