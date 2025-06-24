package main

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	array := []int{
		1, 4, 3, 6, 5,
	}

	var n *Node

	for i := 0; i < len(array); i++ {
		insertOrder(&n, array[i])
	}
	n.printList()
}
