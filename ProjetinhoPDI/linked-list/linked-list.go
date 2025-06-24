/*
estrutura que guarda nós ligados, que seguem as seguintes regras:

 - contem uma sequencia de nós
 - um nó tem uma informação valor e guarda o endereco do proximo nó
 - o primeiro é chamado de head
 - o último nó aponta para null
*/

package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func (ll *LinkedList) insertOrder(valor int) {
	newNode := &Node{data: valor}

	if ll.head == nil || valor < ll.head.data {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil && current.next.data < valor {
		current = current.next
	}

	// [1]->[4]
	// [1]->[3]->[4]-> nil

	newNode.next = current.next
	current.next = newNode
}

func main() {

}

// func (ll *LinkedList) insertAtBeginning(data string) {
// 	newNode := &Node{data: data, next: nil}
// 	ll.head = newNode
// }

// func (ll *LinkedList) insertAtEnd(data string) {
// 	newNode := &Node{data: data, next: nil}

// 	if ll.head == nil {
// 		ll.head = newNode

// 		return
// 	}

// 	current := ll.head

// 	for current.next != nil {
// 		current = current.next
// 	}

// 	current.next = newNode
// }

// func (n *Node) findLength() int {
// 	count := 0
// 	current := n
// 	for current != nil {
// 		count++
// 		current = current.next
// 	}
// 	return count
// }

// func (n *Node) printList() {
// 	current := n
// 	for current != nil {
// 		fmt.Print(current.data, " --> ")
// 		current = current.next
// 	}
// 	fmt.Print("null")
// }

func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " --> ")
		current = current.next
	}
	fmt.Print("null")
}

func (ll *LinkedList) ToSlice() []int {
	s := []int{}
	current := ll.head
	for current.next != nil {
		s = append(s, current.data)
	}
	return s
}
