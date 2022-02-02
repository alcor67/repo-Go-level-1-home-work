package sortarray_test

import (
	"fmt"
	"sort"
	"testing"

	"math/rand"
	"time"

	"github.com/alcor67/repo-Go-level-1-home-work/project-sort-numbers/sortarray"
)

func TestSortSmoke(t *testing.T) {
	arr := []int64{-1, 2, -3, 4}
	direction := "+"
	received, err := sortarray.GetInsertSort(arr, direction)
	if err != nil {
		t.Errorf("arry %+v expected to return no error, but received: %s", arr, err)

	}
	expected := []int64{-3, -1, 2, 4}
	for i := 0; i < len(arr); i++ {
		if received[i] != expected[i] {
			t.Errorf("результат %+v не совпалает с проверкой %+v", received[i], expected[i])
		}
	}
}

func ExampleSortArry() {
	arr := []int64{-1, 2, -3, 4}
	direction := "+"
	received, err := sortarray.GetInsertSort(arr, direction)
	fmt.Println(received, err)
	// Output: [-3 -1 2 4] <nil>
}

func TestSortTableOk(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}
	calls := []call{
		{
			Direction: "+",
			Arr:       []int64{-1, 2, -3, 4},
			Expected:  []int64{-3, -1, 2, 4},
		},
		{
			Direction: "-",
			Arr:       []int64{-1, 2, -3, 4},
			Expected:  []int64{4, 2, -1, -3},
		},
	}

	for _, c := range calls {
		received, err := sortarray.GetInsertSort(c.Arr, c.Direction)
		if err != nil {
			t.Errorf("expected to return no error, but received: %s", err)
		}
		for i := 0; i < len(c.Arr); i++ {
			if c.Expected[i] != received[i] {
				t.Errorf("результат %+v не совпалает с проверкой %+v", c.Expected[i], received[i])
				break
			}
		}
	}
}

func TestSortTableFailure(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
	}
	calls := []call{
		//Not supported operation
		{
			Direction: "a",
			Arr:       []int64{-1, 2, -3, 4},
		},
		{
			Direction: "*",
			Arr:       []int64{-1, 2, -3, 4},
		},
		{
			Direction: "%",
			Arr:       []int64{-1, 2, -3, 4},
		},
	}

	for _, c1 := range calls {

		_, err := sortarray.GetInsertSort(c1.Arr, c1.Direction)

		if err == nil {
			t.Errorf("expected to return error, but received no error")

		}
	}
}

func TestSortGlobal(t *testing.T) {
	type call struct {
		Direction string
		Arr       []int64
		Expected  []int64
	}

	for i := 0; i < 100; i++ {
		randSlice, _ := getRandSlice()

		verifiedRandSliseIncr := make([]int64, len(randSlice))
		copy(verifiedRandSliseIncr, randSlice)

		verifiedRandSliseDecr := make([]int64, len(randSlice))
		copy(verifiedRandSliseDecr, randSlice)

		sort.Slice(verifiedRandSliseIncr, func(i, j int) bool { return verifiedRandSliseIncr[i] < verifiedRandSliseIncr[j] })

		sort.Slice(verifiedRandSliseDecr, func(i, j int) bool { return verifiedRandSliseDecr[i] > verifiedRandSliseDecr[j] })

		calls := []call{

			{
				Direction: "+",
				Arr:       randSlice,
				Expected:  verifiedRandSliseIncr,
			},
			{
				Direction: "-",
				Arr:       randSlice,
				Expected:  verifiedRandSliseDecr,
			},
		}

		for _, c := range calls {
			received, err := sortarray.GetInsertSort(c.Arr, c.Direction)
			if err != nil {
				t.Errorf("expected to return no error, but received: %s", err)
			}
			for i := 0; i < len(c.Arr); i++ {
				if c.Expected[i] != received[i] {
					t.Errorf("результат %+v[%v] не совпалает с проверкой %+v[%v]", c.Expected[i], i, received[i], i)
					break
				}
			}
		}
	}
}

func getRandSlice() ([]int64, error) {
	var randNumber int64
	inputNums := []int64{}
	rand.Seed(getRandSeed())
	cycleLen := rand.Intn(10000)

	for i := 0; i < cycleLen; i++ {

		rand.Seed(getRandSeed())
		randNumberNew := rand.Int63()

		if randNumber == randNumberNew {
			continue
		}
		randNumber = randNumberNew
		randOutPut := (randNumber/1e15 - 5000) * 2
		inputNums = append(inputNums, randOutPut)
	}
	return inputNums, nil
}

func getRandSeed() int64 {
	var randSeed int64 = 0
	timeNano := time.Now().UnixNano()

	randSeedNew := int64(timeNano-timeNano/1e6*1e6) / 10

	if randSeed != randSeedNew {
		randSeed = randSeedNew
	}
	return randSeed
}
