/*Exercicio 1:

Escreva um programa que encontre a soma de todos os numeros de 1 a num. O numero sempre sera um numero inteiro positivo
maior que 0. Sua função so precisa retornar o resultado, o que é mostrado entre parenteses no exemplo abaixo é como
voce chega a esse resultado e não faz parte dele

por exemplo (entrada -> Saida)

2 -> 3 (1 + 2)
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)

*/

package main

import "fmt"

func Sumation(n int) int {
	var sum int

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	n := 10
	fmt.Println("A soma de todos os nmeros de 1 a", n, "é:", Sumation(n))
}
