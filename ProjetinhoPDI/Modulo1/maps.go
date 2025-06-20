package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Isabela"] = 24
	idade["Bento"] = 24
	fmt.Println(idade)
	fmt.Println(idade["Isabela"])
	fmt.Println(idade["Bento"])

	anoNascimento := map[string]int{
		"Isabela": 1999,
		"Bento":   1999,
	}

	fmt.Println(anoNascimento)
	fmt.Println(anoNascimento["Isabela"])
	fmt.Println(anoNascimento["Bento"])

	anoNascimento["golang"] = 2000 // Atualizando o valor de Isabela
	fmt.Println(anoNascimento)
	fmt.Println(anoNascimento["golang"])
}
