package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Print("\n") // \n это перенос строки, или fmt.Println() что то же самое
	}
}
