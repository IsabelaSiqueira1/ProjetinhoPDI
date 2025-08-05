package ficheiro_write

import (
	"fmt"
	"os"
)

var caminho = "ficheiro.txt"
var paises = []string{"Brasil", "Portugal", "Espanha", "Italia", "Alemanha", "França", "Natan", "isa"}

// func main() {
// 	var _, ficheiro = os.Stat(caminho)

// 	if !os.IsNotExist(ficheiro) {
// 		escreverFicheiro(caminho, paises)
// 	}
// }

func escreverFicheiro(caminho string, paises []string) {
	var ficheiro, _ = os.OpenFile(caminho, os.O_WRONLY, 0644)
	defer ficheiro.Close()

	for _, pais := range paises { //a cada iteração escreveramos uma nova linha no ficheiro
		ficheiro.WriteString(pais + "\n")
	}

	ficheiro.Sync()
}

func LerFicheiro() {
	var ficheiro, _ = os.OpenFile(caminho, os.O_RDONLY, 0644)
	defer ficheiro.Close()

	var content = make([]byte, 1024)

	ficheiro.Read(content)
	fmt.Println(string(content))
}
