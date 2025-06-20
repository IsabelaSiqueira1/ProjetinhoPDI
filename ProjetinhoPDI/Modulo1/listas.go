package main

import "fmt"

func main() {
	//array - tamanho fixo
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1]) // Output: Hello World
	fmt.Println(array)              // Output: [Hello World]

	numPrimos := [6]int{2, 3, 5, 7, 11, 13} // array de inteiros
	fmt.Println(numPrimos)                  // Output: [2 3 5 7 11 13]

	//slice - tamanho dinâmico
	//var slice []string
	slice := make([]string, 5) // cria um slice com tamanho 2
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice)

	fmt.Println(len(slice))          // Output: 5 - tamanho do slice
	fmt.Println(cap(slice))          // Output: 5 - capacidade do slice
	slice = append(slice, "Isabela") // adiciona um elemento ao slice
	fmt.Println(slice)               // Output: [Hello World Isabela]

	num := []int{2, 3, 5, 7, 11, 13} // slice de inteiros
	fmt.Println(num)                 // Output: [2 3 5 7 11 13]

	num = append(num, 17) // adiciona um elemento ao slice
	fmt.Println(num)      // Output: [2 3 5 7 11
}
