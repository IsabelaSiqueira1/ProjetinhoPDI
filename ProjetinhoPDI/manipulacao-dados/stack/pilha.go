package pilha

import collection "main.go/interface"

type Stack struct {
	top *Node
}

type Node struct {
	data int
	next *Node
}

func New() collection.Collection {
	return &Stack{}
}

func (s *Stack) Insert(valor int) { //metodo pop
	newNode := &Node{data: valor, next: s.top}
	s.top = newNode
}

func (s *Stack) ForEach(fn func(int)) {
	for current := s.top; current != nil; current = current.next {
		fn(current.data)
	}

}
