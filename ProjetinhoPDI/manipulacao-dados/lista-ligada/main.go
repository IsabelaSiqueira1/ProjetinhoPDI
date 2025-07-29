package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func (ll *LinkedList) insert(valor int) {
	newNode := &Node{data: valor}

	if ll.head == nil || valor < ll.head.data {
		newNode.next = ll.head //o novo no sera o primeiro da lista e o head aponta para o novo no(que antes apontava para outro no que ficou como segundo agora)
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil && current.next.data < valor {
		current = current.next
	}

	newNode.next = current.next
	current.next = newNode
}

func (ll *LinkedList) print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" --> ")
		}
		current = current.next
	}
	fmt.Println(" --> null")
}

func (ll *LinkedList) save(arquivo string) {
	file, _ := os.Create(arquivo)
	defer file.Close()

	current := ll.head
	for current != nil {
		file.WriteString(strconv.Itoa(current.data) + "\n")
		current = current.next
	}
}

func (ll *LinkedList) load(arquivo string) {
	file, _ := os.Open(arquivo)
	defer file.Close()

	content, _ := io.ReadAll(file)
	lines := strings.Split(string(content), "\n")

	ll.head = nil

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if line != "" {
			valor, _ := strconv.Atoi(line)
			ll.insert(valor)
		}
	}
}

func main() {
	lista := &LinkedList{}
	valores := []int{10, 5, 15, 3, 7}

	for _, v := range valores {
		lista.insert(v)
	}

	fmt.Print("lista ligada: ")
	lista.print()

	lista.save("dados.txt")

	novaLista := &LinkedList{}
	novaLista.load("dados.txt")
}
