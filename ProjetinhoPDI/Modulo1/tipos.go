package main

import (
	"fmt"
)

func main() {
	//bool (true/false)%v para formatar tudo(default) %Tmostra o tipo do dado
	fmt.Printf("Type: %T - Value: %v\n", true, true)

	//string - sequencia de bytes
	fmt.Printf("Type: %T - Value: %v\n", "Isabela", "Isabela")

	//int - inteiro
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)

	//float64 - ponto flutuante(flot64/float32) - decimal
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.233)
}
