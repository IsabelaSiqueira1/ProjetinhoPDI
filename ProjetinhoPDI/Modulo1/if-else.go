package main

import "fmt"

func main() {
	numero := 2
	//if {condição} {<ação>}
	if numero == 1 {
		fmt.Println("O valor é igual a 1")
	} else {
		fmt.Println("O valor é diferente de 1")
	}

	if numero == 1 {
		fmt.Println("O valor é igual a 1")
	} else if numero == 2 {
		fmt.Println("O valor é igual a 2")
	} else {
		fmt.Println("O valor é diferente de 1 e 2")
	}

	if 7%2 == 0 {
		fmt.Println("O número 7 é par")
	} else {
		fmt.Println("O número 7 é ímpar")
	}
}
