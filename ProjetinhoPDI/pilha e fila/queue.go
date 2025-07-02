package main

import "fmt"

type Queue struct {
	length int
	front  *ListNode //front -> nil <- rear
	rear   *ListNode
}

type ListNode struct {
	data string
	next *ListNode
}

func (ln *Queue) Length() int {
	return ln.length
}

func (ln *Queue) isEmpty() bool {
	return ln.Length() == 0
}

func (q *Queue) enqueue(data string) { //offer
	temp := &ListNode{data: data}
	if q.isEmpty() {
		q.front = temp
	} else {
		q.rear.next = temp
	}

	q.rear = temp
	q.length++
}

func (q *Queue) dequeue() string { //poll
	if q.isEmpty() {
		panic("is not possible to dequeue an empty queue")
	}

	result := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.length--
	return result
}

func main() {
	q := &Queue{}
	q.enqueue("joao")
	q.enqueue("isabela")
	q.enqueue("natan")
	q.enqueue("golang")

	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
