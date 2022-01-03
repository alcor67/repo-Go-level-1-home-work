package main

import "fmt"

func main() {
	var users = [3]string{2: "Tom", 0: "Alice", 1: "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	var users1 = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users1 {
		fmt.Println(index, value)
	}

	for _, value := range users {
		fmt.Println(value)
	}

	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
