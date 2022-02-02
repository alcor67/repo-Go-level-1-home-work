package fibonacci_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/projectfibonacci/fibonacci"
	"github.com/stretchr/testify/assert"
)

var ExpectedArr = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}
var lenArry = uint(len(ExpectedArr))

func TestFibonacciSmoke(t *testing.T) {

	receivedArr, err := getFibArry(lenArry)
	if err != nil {
		t.Errorf("arry  expected to return no error, but received error")
	}

	var i uint = 0
	for i = 0; i < lenArry; i++ {
		if receivedArr[i] != ExpectedArr[i] {
			t.Errorf("результат %+v не совпалает с проверкой %+v", receivedArr[i], ExpectedArr[i])
		}
	}
}

func ExampleFibonacci() {
	receivedArr, err := getFibArry(lenArry)
	fmt.Println(receivedArr, err)
	// Output: [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181] <nil>
}

func TestTestifyFibonacci(t *testing.T) {

	receivedArr, err := getFibArry(lenArry)

	if err != nil {
		t.Errorf("arry  expected to return no error, but received error")
	}

	assert.Equal(t, ExpectedArr, receivedArr, "they should be equal")
}

func getFibArry(lenArry uint) ([]uint64, error) {

	receivedArr := make([]uint64, lenArry)
	var i uint = 0
	for i = 0; i < lenArry; i++ {

		Expected, err := fibonacci.FibWrap(true)
		if err != nil {
			return receivedArr, errors.New("arry  expected to return no error, but received error")
		}
		receivedArr[i] = Expected(i)
	}
	return receivedArr, nil
}
