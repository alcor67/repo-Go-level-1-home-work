package sortarray

import (
	"bufio"
	"os"
	"strconv"
)

func GetSlice() []int64 {
	inputNums := []int64{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		inputNums = append(inputNums, num)
	}
	return inputNums
}
