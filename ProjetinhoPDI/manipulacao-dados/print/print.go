package print

import (
	"fmt"
)

type Iterador interface {
	ForEach(func(int))
}

func PrintCollection(iterator Iterador) {
	iterator.ForEach(func(v int) {
		fmt.Print(v, " ")
	})
	fmt.Println()
}
