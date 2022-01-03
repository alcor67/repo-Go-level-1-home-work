package main

import "fmt"
/*
после оператора switch мы можем указывать любое выражение, 
которое возвращает значение. Например, операцию сложения:
*/
func main() {
	a := 6
	switch a + 2 {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	}
}
