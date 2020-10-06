package hash_table

import (
	"fmt"
	"strings"
)

const (
	factor       = 1.5
	startSize    = 10
	rehashFactor = 0.8
	startHash    = 6454654
)

var keys = make(map[uint][]string)

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

func (h *HashTable) Size() int {
	return int(h.itemCount)
}

func (h *HashTable) Set(key string, v interface{}) {
	idx := h.hash(key)
	item := h.store[idx]

	if item != nil {
		if item.key == key {
			item.value = v
		} else {
			h.itemCount++
			h.addToNext(item, key, v)
		}
	} else {
		h.itemCount++
		h.store[idx] = &hashItem{
			value: v,
			key:   key,
			next:  nil,
		}
	}

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

func (h *HashTable) GetAllKeys() []string {
	res := make([]string, 0, h.itemCount)

	for _, v := range h.store {
		if v != nil {
			res = append(res, v.key)
		}
	}

	return res
}

func (h *HashTable) Remove(key string) {
	code := h.hash(key)
	item := h.store[code]

	if item != nil {
		if item.key == key {
			if item.next != nil {
				h.store[code] = item.next
			} else {
				h.store[code] = nil
			}
			h.itemCount--
		} else {
			h.removeFromItem(item, key)
		}
	}
}

func (h *HashTable) removeFromItem(item *hashItem, key string) {
	if item.next == nil {
		return
	}

	if item.next.key == key {
		if item.next.next != nil {
			item.next = item.next.next
		} else {
			item.next = nil
		}
		h.itemCount--
	} else {
		h.removeFromItem(item.next, key)
	}
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

func (h *HashTable) addToNext(item *hashItem, key string, value interface{}) {
	if item.next == nil {
		item.next = &hashItem{
			value: value,
			key:   key,
			next:  nil,
		}
	} else if item.next.key == key {
		item.next.value = value
	} else {
		h.addToNext(item.next, key, value)
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
			v.next = nil
			h.addToTmp(tmp, v)

			for next != nil {
				tmpNext := next.next
				next.next = nil
				h.addToTmp(tmp, next)
				next = tmpNext
			}
		}
	}

	h.store = tmp
}

func (h *HashTable) addToTmp(tmp []*hashItem, item *hashItem) {
	code := h.hash(item.key)

	if tmp[code] != nil {
		if tmp[code].key == item.key {
			tmp[code].value = item.value
		} else {
			h.addToNext(tmp[code], item.key, item.value)
		}
	} else {
		tmp[code] = item
	}
}

func (h *HashTable) hash(key string) uint {
	code := h.hashCode(key)

	if keyList, ok := keys[code]; ok {
		add := true
		for _, y := range keyList {
			if y == key {
				add = false
			}
		}

		if add {
			keyList = append(keyList, key)
		}
		keys[code] = keyList
	} else {
		keys[code] = []string{key}
	}

	return code % h.size
}

func (h *HashTable) hashCode(s string) uint {
	hashCode := uint(startHash)

	b := []byte(s)

	for _, v := range b {
		hashCode = hashCode<<5 + uint(v)
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
