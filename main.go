package main

import (
	hashtable "algo-9/hash-table"
	"fmt"
	"math"
	"strconv"
)

var testData = []string{"cat", "tac", "dog", "man", "some", "why", "how", "wow", "wtf", "y", "test", "how do you do", "x", "z"}

func main()  {
	t := hashtable.NewHashTable()

	for i := 0; i < 1400; i++ {
		t.Set(strconv.Itoa(i), i)
	}

	//for i, v := range testData {
	//	t.Set(v, fmt.Sprintf("value %v", i))
	//	//fmt.Println(t.String())
	//	//fmt.Printf("\n==============\n")
	//}

	fmt.Println(t.String())
	fmt.Println(t.Get("355"))
	fmt.Println(math.MaxInt8)
}