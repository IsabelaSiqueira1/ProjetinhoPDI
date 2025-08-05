package fila

import collection "main.go/interface"

type Queue struct {
	front *Node // primeiro da fila
	rear  *Node // último da fila
}

type Node struct {
	data int
	next *Node
}

func New() collection.Collection {
	return &Queue{}
}

func (q *Queue) Insert(valor int) {
	newNode := &Node{data: valor}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) ForEach(fn func(int)) {
	for current := q.front; current != nil; current = current.next {
		fn(current.data)
	}
}
