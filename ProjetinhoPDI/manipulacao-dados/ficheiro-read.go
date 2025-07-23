package main

import (
	"fmt"
	"os"
)

var caminho = "ficheiro.txt"

func main() {
	var _, ficheiro = os.Stat(caminho)

	if !os.IsNotExist(ficheiro) {
		lerFicheiro()
	}
}

func lerFicheiro() {
	var ficheiro, _ = os.OpenFile(caminho, os.O_RDONLY, 0644)
	defer ficheiro.Close()

	var content = make([]byte, 1024)

	ficheiro.Read(content)
	fmt.Println(string(content))
}
