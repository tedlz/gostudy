package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	single()
	test()
}

func single() {
	reward := []int64{1, 2, 3, 2, 3, 1, 3, 1}
	probability := []int64{50, 50, 0, 0, 0, 0, 0, 0}

	rand.Seed(time.Now().UnixNano())
	var lotteryId, lotteryIndex int64
	var sum int64
	type Mix struct {
		Id    int64
		Index int
	}
	var pool = []Mix{}
	for k, v := range reward {
		sum = sum + probability[k]
		for i := 0; i < int(probability[k]); i++ {
			pool = append(pool, Mix{v, k})
		}
	}
	random := rand.Int63n(sum)
	lotteryId = pool[random].Id
	lotteryIndex = int64(pool[random].Index)
	fmt.Println(lotteryIndex, lotteryId)
}

func test() {
	one := 0
	two := 0
	three := 0
	other := 0
	reward := []int64{1, 2, 3}
	probability := []int64{1, 9, 90}

	rand.Seed(time.Now().Unix())
	var lotteryId int64
	var lotteryIndex int64
	for i := 0; i < 10000; i++ {
		type Mix struct {
			Id    int64
			Index int
		}
		var pool = []Mix{}
		var sum int64
		for k, v := range reward {
			sum += probability[k]
			for j := 0; j < int(probability[k]); j++ {
				pool = append(pool, Mix{v, k})
			}
		}
		random := rand.Int63n(sum)
		lotteryId = pool[random].Id
		lotteryIndex = int64(pool[random].Index)
		if lotteryIndex == 0 {
			one += 1
		} else if lotteryIndex == 1 {
			two += 1
		} else if lotteryIndex == 2 {
			three += 1
		} else {
			other += 1
		}
	}
	fmt.Println(one, two, three, other, lotteryId, lotteryIndex)
}
