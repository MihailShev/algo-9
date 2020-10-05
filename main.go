package main

import (
	hashtable "algo-9/hash-table"
	testdata "algo-9/test-data"
	"algo-9/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := hashtable.NewHashTable()

	//fmt.Println(text)

	keys := strings.Split(testdata.TestText, " ")
	fmt.Println(len(keys))
	orderDataSet := utils.FillArray(1, 10_000_000)
	valuesToSearch := utils.GetRandomValueList(orderDataSet, len(orderDataSet)/10)
	t1 := hashtable.NewHashTable()
	t2 := make(map[string]int)

	test(func() {
		for _, v := range orderDataSet {
			t1.Set(strconv.Itoa(v), v)
		}
	})

	test(func() {
		for _, v := range valuesToSearch {
			t1.Get(strconv.Itoa(v))
		}
	})

	test(func() {
		for _, v := range orderDataSet {
			t2[strconv.Itoa(v)] = v
		}
	})

	test(func() {
		for _, v := range valuesToSearch {
			_, _ = t2[strconv.Itoa(v)]
		}
	})

	for _, key := range keys {
		v := t.Get(key)
		if v == nil {
			t.Set(key, 1)
		} else {
			count := v.(int)
			count++
			t.Set(key, count)
		}
	}

	fmt.Println(t.String())

	allKeys := t.GetAllKeys()

	maxCount := 0
	key := ""

	for _, k := range allKeys {
		if v := t.Get(k); v != nil {
			intCount := v.(int)
			if intCount > maxCount {
				maxCount = intCount
				key = k
			}
		}
	}

	fmt.Println("key", key, "count", maxCount)
	fmt.Println(t.Get("Он"))
}

func test(run func()) {
	start := time.Now()
	run()
	stop := time.Since(start)
	fmt.Printf("Execution time: %s\n\n", stop)
}
