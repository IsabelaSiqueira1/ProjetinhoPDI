package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13) // função soma recebe dois parametros inteiros e retorna um inteiro
	fmt.Println(somaDosValores)

	sub := substracao(10, 5) // função subtracao recebe dois parametros inteiros e retorna um inteiro
	fmt.Println(sub)

	nome := printaNome("Isabela") // função printaNome recebe um parametro string e retorna uma string
	fmt.Println(nome)
}

func soma(x int, y int) int {
	return x + y
}

func substracao(x int, y int) int {
	return x - y
}

func printaNome(nome string) string {
	return nome
}
