package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("O valor de i é:", i)
	}
	fmt.Println("A soma dos números de 0 a 9 é:", sum)

	//loop infinito
	// for {
	// 	fmt.Println("Golandingggggggg!!!!!")
	// }

	//for range
	frutas := []string{"maçã", "banana", "laranja"}
	for _, fruta := range frutas {
		fmt.Println("Fruta:", fruta)
	}
}
