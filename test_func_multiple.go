package main

import "fmt"

func main() {
	//после списка параметров указывается в скобках список типов возвращаемых значений
	//var age, name = add(4, 5, "Tom", "Simpson")
	//Альтернативный способ передачи переменным результатов функции
	myAge, myName := add(4, 5, "Tom", "Simpson")
	fmt.Println("Age: ", myAge)       // 9
	fmt.Println("Fullname: ", myName) // Tom Simpson
}
func add(x, y int, firstName, lastName string) (age int, fullName string) {
	age = x*10 + y
	fullName = firstName + " " + lastName
	//после оператора return располагаются через запятую все возвращаемые значения
	return
}
