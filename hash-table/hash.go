package hash_table

import (
	"fmt"
	"strings"
)

const (
	factor       = 1.5
	startSize    = 10
	rehashFactor = 0.8
	startHash = 6454654
)

type HashTable struct {
	itemCount  uint
	size       uint
	store      []*hashItem
	rehashSize uint
}

type hashItem struct {
	value interface{}
	key   string
	next  *hashItem
}

func NewHashTable() HashTable {
	rehashSize := startSize * rehashFactor

	h := HashTable{
		itemCount:  0,
		size:       startSize,
		rehashSize: uint(rehashSize),
		store:      make([]*hashItem, startSize),
	}

	return h
}

func (h *HashTable) Set(key string, v interface{}) {
	idx := h.hash(key)
	item := h.store[idx]

	if item != nil {
		if item.key == key {
			item.value = v
		} else {
			h.addToItem(item, key, v)
		}
	} else {
		h.store[idx] = &hashItem{
			value: v,
			key:   key,
			next:  nil,
		}
	}

	h.itemCount++

	if h.isNeedRehash() {
		h.rehash()
	}
}

func (h *HashTable) Get(key string) interface{} {
	code := h.hash(key)
	item := h.store[code]

	if item == nil {
		return nil
	}

	return searchInItem(item, key)
}

func searchInItem(item *hashItem, key string) interface{} {
	if item.key == key {
		return item.value
	} else {
		if item.next != nil {
			return searchInItem(item.next, key)
		}
	}

	return nil
}

func (h *HashTable) addToItem(item *hashItem, key string, value interface{}) {
	if item.next == nil {
		item.next = &hashItem{
			value: value,
			key:   key,
			next:  nil,
		}
	} else {
		h.addToItem(item.next, key, value)
	}
}

func (h *HashTable) isNeedRehash() bool {
	return h.itemCount > h.rehashSize
}

func (h *HashTable) rehash() {
	size := float64(h.size) * factor
	rehashSize := size * rehashFactor

	h.size = uint(size)
	h.rehashSize = uint(rehashSize)

	tmp := make([]*hashItem, h.size)
	var next *hashItem

	for _, v := range h.store {

		if v != nil {
			next = v.next
			h.addToTmp(tmp, v)

			for next != nil {
				tmpNext := next.next
				h.addToTmp(tmp, next)
				next = tmpNext
			}
		}
	}

	h.store = tmp
}

func (h *HashTable) addToTmp(tmp []*hashItem, item *hashItem) {
	code := h.hash(item.key)
	item.next = nil

	if tmp[code] != nil {
		h.addToItem(tmp[code], item.key, item.value)
	} else {
		tmp[code] = item
	}
}

func (h *HashTable) hash(key string) uint {
	code := h.hashCode(key)
	return code % h.size
}

func (h *HashTable) hashCode(s string) uint {
	hashCode := uint(startHash)

	b := []byte(s)

	for _, v := range b {
		hashCode = hashCode << 5 + uint(v)
	}

	return hashCode
}

func (h *HashTable) String() string {
	str := strings.Builder{}

	for i, v := range h.store {
		str.WriteString(fmt.Sprintf("%v ", i))
		if v != nil {

			str.WriteString(fmt.Sprintf("key: \"%v\": value: \"%v\" -> ", v.key, v.value))

			next := v.next

			for next != nil {
				str.WriteString(fmt.Sprintf("key \"%v\": value \"%v\" -> ", next.key, next.value))
				next = next.next
			}
			str.WriteString("\n")
		} else {
			str.WriteString(fmt.Sprintf("%v\n", v))
		}
	}

	return str.String()
}
