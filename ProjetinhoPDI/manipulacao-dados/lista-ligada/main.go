package main

import (
	"bufio"
	"fmt"
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
	//fmt.Printf("Inserido %d na lista ligada\n", valor)
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
	file, _ := os.OpenFile(arquivo, os.O_CREATE|os.O_WRONLY, 0644)
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

	content := bufio.NewScanner(file) // readAll() não utilizar alocação de memoria desnecessária
	//lines := strings.Split(string(content), "\n")
	//ll.head = nil somente se quiser uma lista vazia antes de carregar os dados

	for content.Scan() {
		line := strings.TrimSpace(content.Text())

		if line != "" {
			valor, _ := strconv.Atoi(line)
			//fmt.Printf("valor do load %d \n", valor)
			ll.insert(valor)
		}
	}
}

func main() {
	lista := &LinkedList{}
	valores := []int{10, 5, 15, 3, 7, 50, 11, 12}

	lista.print()

	for _, v := range valores {
		lista.insert(v)
	}

	lista.print()

	lista.load("dados.txt")
	fmt.Print("load lista ligada: ")
	lista.print()

	fmt.Print("lista ligada: ")
	lista.print()

	lista.save("dados.txt")

	fmt.Println("Lista ligada salva em dados.txt")

}
