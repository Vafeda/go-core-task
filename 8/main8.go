package main

import (
	"fmt"
)

const (
	BucketSize = 16
)

type HashTable struct {
	array []*bucket
}

type bucket struct {
	bucket []Node
	next   *bucket
}

type Node struct {
	key   string
	value int
}

func New(arraySize uint) *HashTable {
	array := make([]*bucket, arraySize)

	for i := range array {
		array[i] = &bucket{}
	}

	return &HashTable{
		array: array,
	}
}

func (h *HashTable) Add(key string, value int) {
	index := h.hash(key)

	h.array[index].add(key, value)
}

func (b *bucket) add(key string, value int) {
	current := b

	for {
		for i := range current.bucket {
			if current.bucket[i].key == key {
				current.bucket[i].value = value
				return
			}
		}

		if len(current.bucket) < BucketSize {
			current.bucket = append(current.bucket, Node{key, value})
			return
		}

		if current.next == nil {
			current.next = new(bucket)
		}
		current = current.next
	}
}

func (h *HashTable) Copy() map[string]int {
	copyHashTable := make(map[string]int)

	for _, a := range h.array {
		current := a
		for {
			for _, b := range current.bucket {
				copyHashTable[b.key] = b.value
			}

			if current.next == nil {
				break
			}

			current = current.next
		}
	}

	return copyHashTable
}

func (h *HashTable) Remove(key string) {
	index := h.hash(key)
	h.array[index].remove(key)
}

func (b *bucket) remove(key string) {
	current := b

	for {
		for i := range current.bucket {
			if current.bucket[i].key == key {
				current.bucket = append(current.bucket[:i], current.bucket[i+1:]...)
				return
			}
		}

		if len(current.bucket) < BucketSize {
			return
		}

		if current.next == nil {
			return
		}

		current = current.next
	}
}

func (h *HashTable) Exists(key string) bool {
	index := h.hash(key)
	return h.array[index].exists(key)
}

func (b *bucket) exists(key string) bool {
	current := b

	for {
		for i := range current.bucket {
			if current.bucket[i].key == key {
				return true
			}
		}

		if len(current.bucket) < BucketSize {
			return false
		}

		if current.next == nil {
			return false
		}

		current = current.next
	}
}

func (h *HashTable) Get(key string) (int, bool) {
	index := h.hash(key)
	return h.array[index].get(key)
}

func (b *bucket) get(key string) (int, bool) {
	current := b

	for {
		for i := range current.bucket {
			if current.bucket[i].key == key {
				return current.bucket[i].value, true
			}
		}

		if len(current.bucket) < BucketSize {
			return 0, false
		}

		if current.next == nil {
			return 0, false
		}

		current = current.next
	}
}

func (h *HashTable) hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % len(h.array)
}

func main() {
	hashTable := New(7)

	fmt.Println("Hash Table Created")
	for i := 0; i < 8; i++ {
		hashTable.Add(fmt.Sprintf("%d", i), i)
		hashTable.Add(fmt.Sprintf("%d", i), i+1)
	}

	for i, v := range hashTable.array {
		fmt.Printf("Bucket %d: [", i)
		for _, b := range v.bucket {
			fmt.Printf("[%s,%d], ", b.key, b.value)
		}
		fmt.Println("]")
	}

	fmt.Println(hashTable.Copy())
}
