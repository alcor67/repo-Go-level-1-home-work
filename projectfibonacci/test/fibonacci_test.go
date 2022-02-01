package fibonacci_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/projectfibonacci/fibonacci"
	"github.com/stretchr/testify/assert"
)

var fibNumAarr = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}
var lenArry = uint(len(fibNumAarr))

func TestFibonacciSmoke(t *testing.T) {

	ExpectedArr, err := getFibArry(lenArry)
	if err != nil {
		t.Errorf("arry  expected to return no error, but received error")
	}

	/*
		make([]uint64, lenArry)
		var i uint = 0
		for i = 0; i < lenArry; i++ {

			Expected, err := fibonacci.FibWrap(true)
			if err != nil {
				t.Errorf("arry %+v expected to return no error, but received: %s", Expected, err)
			}
			fibNumAarr[i] = Expected(i)
		}
	*/
	var i uint = 0
	for i = 0; i < lenArry; i++ {
		if ExpectedArr[i] != fibNumAarr[i] {
			t.Errorf("результат %+v не совпалает с проверкой %+v", ExpectedArr[i], fibNumAarr[i])
		}
	}
}

func ExampleFibonacci() {
	ExpectedArr, err := getFibArry(lenArry)
	fmt.Println(ExpectedArr, err)
	// Output: [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181] <nil>
}

func TestTestifyFibonacci(t *testing.T) {

	ExpectedArr, err := getFibArry(lenArry)

	if err != nil {
		t.Errorf("arry  expected to return no error, but received error")
	}

	assert.Equal(t, ExpectedArr, fibNumAarr, "they should be equal")
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
