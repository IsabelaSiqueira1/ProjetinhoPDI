package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	front *Node // primeiro da fila
	rear  *Node // último da fila
}

type Node struct {
	data int
	next *Node
}

func (q *Queue) enqueue(valor int) {
	newNode := &Node{data: valor}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

var queueEmptyErr = errors.New("Fila vazia")

func (q *Queue) dequeue() (int, error) {
	if q.front == nil {
		return -1, queueEmptyErr
	}

	valor := q.front.data
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	return valor, nil
}

func (q *Queue) print() {
	if q.front == nil {
		fmt.Println("Fila vazia")
		return
	}

	current := q.front
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" -> Fim")
}

func (q *Queue) save(arquivo string) {
	file, _ := os.Create(arquivo)
	defer file.Close()

	current := q.front
	for current != nil {
		file.WriteString(strconv.Itoa(current.data) + "\n")
		current = current.next
	}
}

func (q *Queue) load(arquivo string) {
	file, _ := os.Open(arquivo)
	defer file.Close()

	content, _ := io.ReadAll(file)
	lines := strings.Split(string(content), "\n")

	q.front = nil
	q.rear = nil

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			valor, _ := strconv.Atoi(line)
			q.enqueue(valor)
		}
	}
}

func main() {
	fila := &Queue{}
	valores := []int{3, 10, 15, 3, 7, 2, 1, 8, 10, 12, 20, 24, 55}

	for _, v := range valores {
		fila.enqueue(v)
	}

	fmt.Print("Fila: ")
	fila.print()

	fila.save("dados_fila.txt")

	novaFila := &Queue{}
	novaFila.load("dados_fila.txt")

	// remove := fila.dequeue()
	// fmt.Println("Removido da fila:", remove)

	// fmt.Print("Fila após remoção: ")
	// fila.print()
}
