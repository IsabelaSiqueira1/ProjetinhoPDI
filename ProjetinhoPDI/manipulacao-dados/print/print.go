package print

import (
	"fmt"

	collection "main.go/interface"
)

func PrintCollection(collection collection.Collection) {
	collection.ForEach(func(v int) {
		fmt.Print(v, " ")
	})
	fmt.Println()
}
