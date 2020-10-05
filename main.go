package main

import (
	hashtable "algo-9/hash-table"
	testdata "algo-9/test-data"
	"fmt"
	"strings"
)

func main() {
	t := hashtable.NewHashTable()

	//fmt.Println(text)

	keys := strings.Split(testdata.TestText, " ")
	fmt.Println(len(keys))

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
