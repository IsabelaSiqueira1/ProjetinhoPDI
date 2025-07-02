// package main

// import "fmt"

// type StackLinkedList struct {
// 	Top    *StackNode
// 	Length int
// }

// type StackNode struct {
// 	data int
// 	Down *StackNode
// }

// func (sl *StackLinkedList) Push(data int) {
// 	newNode := &StackNode{data: data}
// 	newNode.Down = sl.Top
// 	sl.Top = newNode
// }

// func (sl *StackLinkedList) Pop() int {
// 	result := sl.Top.data
// 	sl.Top = sl.Top.Down
// 	sl.Length--
// 	return result
// }

// func main() {
// 	stack := &StackLinkedList{}
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)

// 	fmt.Println(stack.Pop())
// }
