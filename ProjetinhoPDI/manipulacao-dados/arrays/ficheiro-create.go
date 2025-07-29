package main

import (
	"fmt"
	"os"
)

var caminho = "ficheiro.txt"
var paises = [...]string{"Brasil", "Portugal", "Espanha", "Italia", "Alemanha", "França"}

func main() {
	// if _, err := os.Stat(caminho); os.IsNotExist(err) {
	// 	f, err := os.Create(caminho)
	// 	if err != nil {
	// 		fmt.Println("Erro criando arquivo:", err)
	// 		return
	// 	}
	// 	defer f.Close()
	// 	fmt.Println("Arquivo criado:", caminho)
	// 	escreverFicheiro(paises, f)
	// }

	//fmt.Print("ja existe o arquivo")

	var ficheiro, _ = os.OpenFile(caminho, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer ficheiro.Close()

	for _, pais := range paises {
		_, err := ficheiro.WriteString(pais + "\n")
		if err != nil {
			fmt.Println("Erro escrevendo no arquivo:", err)
			return
		}
	}

	ficheiro.Sync()
}

func escreverFicheiro(paises [6]string, f *os.File) {
	// _, err := os.OpenFile(caminho, os.O_APPEND|os.O_WRONLY, 0644) // write-only
	// if err != nil {
	// 	fmt.Println("Erro abrindo arquivo para escrita:", err)
	// 	return
	// }

	for _, pais := range paises {
		f.WriteString(pais + "\n")
	}

	f.Sync()
}
