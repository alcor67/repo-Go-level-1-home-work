package main

import (
	"fmt"
	"time"
)

func main() {
	var n uint
	fmt.Print("введите число членов последовательности Фибоначчи: \n")
	fmt.Scanln(&n)

	//присваиваем значение переменной getFibNum функцию обертку для замыкания,
	//которая будет возвращать функцию считающую число Фибоначчи
	getFibNumCahe := fibWrap(true) // с кешем
	getFibNumNoCahe := fibWrap()   // без кеша

	//ставим замер времении
	now := time.Now()
	fmt.Println("число Фибоначчи: ", getFibNumCahe(n))
	fmt.Println("количество итераций с кешем: ", i)
	//выводим замер времении
	fmt.Printf("время рассчета с кешем: %d ms\n", time.Since(now).Milliseconds())
	i = 0 // обнуляем счетчик итераций
	now = time.Now()
	fmt.Println("число Фибоначчи: ", getFibNumNoCahe(n))
	fmt.Println("количество итераций без кеша: ", i)
	fmt.Printf("время рассчета без кеша: %d ms\n", time.Since(now).Milliseconds())
}

var i uint = 0

//создаем обертку для замыкания, которая будет возвращать функцию считающую число Фибоначчи
//входящие аргументы и именной выходной параметр можно не указывать, только их типы
func fibWrap(caheArg ...bool) func(uint) uint64 {
	// объявление и инициализация мапы через var для создания кеша
	//а можно и так
	//var mapFibNumCache map[uint]uint64 = make(map[uint]uint64)
	var mapFibNumCache = make(map[uint]uint64)
	//объявляем переменную через сигнатуру функции,
	//сигнатура должна совпадать с функцией считающей число Фибоначчи
	var fibNum func(uint) uint64

	//используем именной выходной параметр result
	//добавляем аргумент для подключения кеша caheArg
	//... значит множество параметров, в т. числе и ничего, т.е. без этого параметра
	//если ничего (по умолчанию), то len(caheArg) = 0 и caheArg[0] = false
	//т.е. по умолчанию участок кода с кешем отключен
	//присваиваем значение переменной fibNum анонимную функцию, считающую число Фибоначчи
	fibNum = func(n uint) (result uint64) {
		i++
		if len(caheArg) > 0 && caheArg[0] == true {
			//сначала ищем результат в кеше, если есть, то возвращаем.
			//здесь result  это значение мапы, ok это логический true/false признак существования элемента мапы
			//т.е. есть ли наш результат result в мапе?
			//заглядываем в мапу по ключу n
			if result, ok := mapFibNumCache[n]; ok {
				return result
			}
			//применяем отложенную функцию, которая кладет в мапу (которая кеш) элемент result
			defer func() {
				mapFibNumCache[n] = result
			}()
		}

		if n <= 1 {
			return uint64(n)
		} else {
			//... значит множество параметров, в т. числе и ничего, т.е. без этого параметра
			return fibNum(n-1) + fibNum(n-2)
		}
	}
	return fibNum
}
