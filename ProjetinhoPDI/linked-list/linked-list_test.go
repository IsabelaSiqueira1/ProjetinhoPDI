package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	array := []int{
		1, 4, 3, 6, 5, 0, 10,
	}

	list := &LinkedList{head: nil}

	for i := 0; i < len(array); i++ {
		list.insertOrder(array[i])
	}

	list.printList()

	got := list.ToSlice()
	t.Logf("%v\n", array)
	t.Logf("%v\n", got)
}
