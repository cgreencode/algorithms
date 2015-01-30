// Implementation of a Hash Table with Separate Chaining
// with linked lists using Horner's hash function
// http://en.wikipedia.org/wiki/Hash_table#Separate_chaining_with_linked_lists
package ht

import (
	"errors"
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/linked-list"
)

type HashTable struct {
	Table    map[int]*list.List
	Size     int
	Capacity int
}

type item struct {
	key   string
	value string
}

func New(cap int) *HashTable {
	table := make(map[int]*list.List, cap)
	return &HashTable{Table: table, Size: 0, Capacity: cap}
}

func (ht *HashTable) Get(key string) (string, error) {
	index := ht.position(key)
	item, err := ht.find(index, key)

	if item == nil {
		return "", errors.New("Not Found")
	}

	fmt.Println(item)
	return item.value, err
}

func (ht *HashTable) Put(key, value string) {
	index := ht.position(key)

	if ht.Table[index] == nil {
		ht.Table[index] = list.NewList()
	}

	item := &item{key: key, value: value}

	a, err := ht.find(index, key)
	if err != nil {
		// The key doesn't exist in HashTable
		ht.Table[index].Append(list.NewNode(item))
		ht.Size++
	} else {
		// The key exists so we overwrite its value
		a.value = value
		fmt.Println(a)
	}
}

func (ht *HashTable) position(s string) int {
	return hashCode(s) % ht.Capacity
}

func (ht *HashTable) find(i int, key string) (*item, error) {
	l := ht.Table[i]
	index, err1 := l.Find(&list.Node{Value: i})
	value, err2 := l.Get(index)

	if err1 != nil || err2 != nil {
		return nil, errors.New("Not Found")
	}

	return value.Value.(*item), nil
}