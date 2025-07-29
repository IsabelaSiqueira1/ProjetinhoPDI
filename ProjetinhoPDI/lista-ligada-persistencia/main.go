package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Estrutura da lista ligada
type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

// Insere valor mantendo ordem
func (ll *LinkedList) insert(valor int) {
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

	newNode.next = current.next
	current.next = newNode
}

func (ll *LinkedList) print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println(" -> null")
}

func (ll *LinkedList) save(arquivo string) {
	file, _ := os.Create(arquivo)
	defer file.Close()

	// Escreve valores diretamente, sem contar primeiro
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

	// Limpa lista atual
	ll.head = nil

	// Lê todos os valores do arquivo, começando da primeira linha
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			valor, _ := strconv.Atoi(line)
			ll.insert(valor)
		}
	}
}

func main() {
	fmt.Println("=== Lista Ligada com Persistência (SEM LIMITE) ===")

	lista := &LinkedList{}
	valores := []int{10, 5, 15, 3, 7}

	fmt.Println("Adicionando valores:", valores)
	for _, v := range valores {
		lista.insert(v)
	}

	fmt.Print("Lista: ")
	lista.print()

	lista.save("dados.txt")
	fmt.Println("Lista salva em 'dados.txt'")

	novaLista := &LinkedList{}
	novaLista.load("dados.txt")

	fmt.Print("Lista carregada: ")
	novaLista.print()
}
