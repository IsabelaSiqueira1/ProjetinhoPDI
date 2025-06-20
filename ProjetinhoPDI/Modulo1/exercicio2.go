/*Exercicio 2:

Nesta tarefa simples, voce recebe um numero e deve torna-lo negativo. Mas talvez o numero ja seja negativo?

Exemplos

makeNegative(1) ->  return -1
makeNegative(-5) -> return -5
makeNegative(0) -> return 0

Nota: O numero ja pode ser negativo, e nesse caso, nenhuma alteração é necessaria.
Zero(0) não é verificado para nenhum sinal especifico. Zeros negativos não fazem sentido matematico.

*/

package main

import "fmt"

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	}
	negativeNum := x * (-1)
	return negativeNum
}

func main() {
	fmt.Println(MakeNegative(1))
	fmt.Println(MakeNegative(-5))
	fmt.Println(MakeNegative(0))
}
