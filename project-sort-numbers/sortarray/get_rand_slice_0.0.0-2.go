package sortarray

import (
	"math/rand"
	"time"
)

func GetRandSlice() ([]int64, error) {
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
	//fmt.Println(len(inputNums))
	//fmt.Println(inputNums)
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
