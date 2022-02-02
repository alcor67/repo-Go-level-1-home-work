package fibonacci_test

import (
	"errors"
	"testing"

	"github.com/alcor67/repo-Go-level-1-home-work/projectfibonacci/fibonacci"
	"github.com/stretchr/testify/assert"
)

var fibNumAarr = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}
var lenArry = uint(len(fibNumAarr))

func TestTestifyFibonacci(t *testing.T) {

	ReceivedArr, err := getFibArry(lenArry)
	if err != nil {
		t.Errorf("arry  expected to return no error, but received error")
	}

	assert.Equal(t, fibNumAarr, ReceivedArr, "they should be equal")
}

func getFibArry(lenArry uint) ([]uint64, error) {

	receivedArr := make([]uint64, lenArry)
	var i uint = 0
	for i = 0; i < lenArry; i++ {

		received, err := fibonacci.FibWrap(true)
		if err != nil {
			return receivedArr, errors.New("arry  expected to return no error, but received error")
		}
		receivedArr[i] = received(i)
	}
	return receivedArr, nil
}
