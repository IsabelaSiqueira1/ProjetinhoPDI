package lista

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Len() int {
	count := 0
	current := ll.head

	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (ll *LinkedList) Insert(valor int) {
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

func (ll *LinkedList) ForEach(fn func(int)) {
	for current := ll.head; current != nil; current = current.next {
		fn(current.data)
	}
}
