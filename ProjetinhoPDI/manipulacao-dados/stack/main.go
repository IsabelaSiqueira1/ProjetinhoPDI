package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	top *Node
}

type Node struct {
	data int
	next *Node
}

func (s *Stack) push(valor int) {
	newNode := &Node{data: valor}
	newNode.next = s.top
	s.top = newNode
}

func (s *Stack) pop() int {
	if s.top == nil {
		return -1 // pilha vazia
	}

	valor := s.top.data
	s.top = s.top.next // topo agora é o próximo
	return valor
}

func (s *Stack) print() {
	if s.top == nil {
		fmt.Println("Pilha vazia")
		return
	}

	current := s.top
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" --> ")
		}
		current = current.next
	}
	fmt.Println(" --> Base")
}

func (s *Stack) save(arquivo string) {
	file, _ := os.Create(arquivo)
	defer file.Close()

	current := s.top
	for current != nil {
		file.WriteString(strconv.Itoa(current.data) + "\n")
		current = current.next
	}
}

func (s *Stack) load(arquivo string) {
	file, _ := os.Open(arquivo)
	defer file.Close()

	content, _ := io.ReadAll(file)
	lines := strings.Split(string(content), "\n")

	s.top = nil

	valores := []int{}
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			valor, _ := strconv.Atoi(line)
			valores = append(valores, valor)
		}
	}

	for i := len(valores) - 1; i >= 0; i-- {
		s.push(valores[i])
	}
}

func main() {
	pilha := &Stack{}
	valores := []int{10, 5, 15, 3, 7, 4, 6, 10, 3, 2}

	for _, v := range valores {
		pilha.push(v)
	}

	fmt.Print("Pilha: ")
	pilha.print()

	pilha.save("dados_pilha.txt")

	novaPilha := &Stack{}
	novaPilha.load("dados_pilha.txt")
}
