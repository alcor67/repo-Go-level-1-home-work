package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/alcor67/repo-Go-level-1-home-work/projectfibonacci/fibonacci"
)

//var i uint = 0

func main() {
	var n uint
	var nInp string
	fmt.Print("введите число членов последовательности Фибоначчи: \n")
	fmt.Scanln(&nInp)
	n1, err := strconv.ParseUint(nInp, 10, 64)

	if err != nil {
		fmt.Println(err.Error())
		return
	} else if n1 == 0 {
		fmt.Println("Введено ноль")
		return
	}
	n = uint(n1)

	getFibNumCahe, err := fibonacci.FibWrap(true)
	if err != nil {
		fmt.Println("arry %+v expected to return no error, but received: %s", getFibNumCahe, err)
		os.Exit(1)
	}

	fmt.Println("число Фибоначчи: ", getFibNumCahe(n-1))
	ExpectedArr, err := getFibArry(n)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Ряд Фибоначчи: ",ExpectedArr)
}

func getFibArry(lenArry uint) ([]uint64, error) {

	ExpectedArr := make([]uint64, lenArry)
	var i uint = 0
	for i = 0; i < lenArry; i++ {

		Expected, err := fibonacci.FibWrap(true)
		if err != nil {
			return ExpectedArr, errors.New("arry  expected to return no error, but received error")
		}
		ExpectedArr[i] = Expected(i)
	}
	return ExpectedArr, nil
}
