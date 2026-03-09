package linkedlist

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	len  int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Insert(value T) {
	newNode := &node[T]{value: value}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.len++
		return
	}

	l.tail.next = newNode
	l.tail = newNode
	l.len++
}

func (l *LinkedList[T]) Find(match func(T) bool) (T, bool) {
	for current := l.head; current != nil; current = current.next {
		if match(current.value) {
			return current.value, true
		}
	}

	var zero T
	return zero, false
}

func (l *LinkedList[T]) Update(match func(T) bool, update func(*T)) (T, bool) {
	for current := l.head; current != nil; current = current.next {
		if match(current.value) {
			update(&current.value)
			return current.value, true
		}
	}

	var zero T
	return zero, false
}

func (l *LinkedList[T]) ForEach(fn func(T)) {
	for current := l.head; current != nil; current = current.next {
		fn(current.value)
	}
}

func (l *LinkedList[T]) Len() int {
	return l.len
}
