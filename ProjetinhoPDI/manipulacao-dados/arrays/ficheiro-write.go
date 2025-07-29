package main

import "os"

var caminho = "ficheiro.txt"
var paises = []string{"Brasil", "Portugal", "Espanha", "Italia", "Alemanha", "França", "Natan", "isa"}

func main() {
	var _, ficheiro = os.Stat(caminho)

	if !os.IsNotExist(ficheiro) {
		escreverFicheiro(caminho, paises)
	}
}

func escreverFicheiro(caminho string, paises []string) {
	var ficheiro, _ = os.OpenFile(caminho, os.O_WRONLY, 0644)
	defer ficheiro.Close()

	for _, pais := range paises { //a cada iteração escreveramos uma nova linha no ficheiro
		ficheiro.WriteString(pais + "\n")
	}

	ficheiro.Sync()
}
