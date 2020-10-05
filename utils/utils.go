package utils

import (
	"math/rand"
	"time"
)

func FillArray(from, to int) []int {
	res := make([]int, 0)

	for i := from; i <= to; i++ {
		res = append(res, i)
	}

	return res
}

func FillArrayUniqRandom(from, to int) []int {
	i := from
	res := make([]int, 0)
	for i <= to {
		res = append(res, i)
		i++
	}

	return mix(res)
}

func GetRandomValueList(arr []int, amount int) []int {
	res := make([]int, 0, amount)
	max := len(arr) - 1
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < amount; i++ {
		res = append(res, arr[getRandom(0, max)])
	}

	return res
}

func mix(arr []int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	size := len(arr)
	for i := size - 1; i >= 0; i-- {
		j := getRandom(0, size)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func getRandom(min, max int) int {
	return min + rand.Intn(max-min)
}
