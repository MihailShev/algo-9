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

	fmt.Println("Test the hash table by counting the number of words in the text")
	t := hashtable.NewHashTable()
	words := strings.Split(testdata.TestText, " ")

	for _, key := range words {
		v := t.Get(key)
		if v == nil {
			t.Set(key, 1)
		} else {
			count := v.(int)
			count++
			t.Set(key, count)
		}
	}

	allKeys := t.GetAllKeys()

	for _, key := range allKeys {
		fmt.Println("word:", key, "count:", t.Get(key))
	}

	fmt.Println("\nComparison the hash table with native implementation")

	orderDataSet := utils.FillArray(1, 10_000_000)
	valuesToSearch := utils.GetRandomValueList(orderDataSet, len(orderDataSet)/10)
	valuesToRemove := utils.GetRandomValueList(orderDataSet, len(orderDataSet)/10)
	fmt.Println(len(valuesToRemove))
	t1 := hashtable.NewHashTable()
	t2 := make(map[string]int)

	fmt.Println("\nAdd 10 000 000 elements to the hash table")
	test(func() {
		for _, v := range orderDataSet {
			t1.Set(strconv.Itoa(v), v)
		}
	})

	fmt.Println("\nAdd 10 000 000 elements to the native map")
	test(func() {
		for _, v := range orderDataSet {
			t2[strconv.Itoa(v)] = v
		}
	})

	fmt.Println("\nGet from the hash table random 1 000 000 elements")
	test(func() {
		for _, v := range valuesToSearch {
			t1.Get(strconv.Itoa(v))
		}
	})

	fmt.Println("\nGet from the native map random 1 000 000 elements")
	test(func() {
		for _, v := range valuesToSearch {
			_, _ = t2[strconv.Itoa(v)]
		}
	})

	fmt.Println("\nRemove from the hash table random 1 000 000 elements")
	test(func() {
		for _, v := range valuesToRemove {
			t1.Remove(strconv.Itoa(v))
		}
	})

	fmt.Println("\nRemove from the  native map random 1 000 000 elements")
	test(func() {
		for _, v := range valuesToRemove {
			delete(t2, strconv.Itoa(v))
		}
	})

	fmt.Println("Finish, press any key")
	_, _ = fmt.Scanf(" ")
}

func test(run func()) {
	start := time.Now()
	run()
	stop := time.Since(start)
	fmt.Printf("Execution time: %s\n\n", stop)
}
