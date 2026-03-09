package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	//"log"
	//"net/http"
	"os"

	leituraescrita "github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/leitura-escrita"

	lista "github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/lista-ligada"
	print "github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/print"
	// stack "github.com/IsabelaSiqueira1/ProjetinhoPDI/ProjetinhoPDI/manipulacao-dados/stack"
)

func processData(name string, struct_data leituraescrita.Collection, reader io.Reader, writer io.Writer) {
	fmt.Println("Estrutura:", name)
	leituraescrita.Read(reader, struct_data)

	fmt.Print("conteúdo da estrutura: ")
	print.PrintCollection(struct_data)

	leituraescrita.Write(writer, struct_data)
	fmt.Println()
}

func main() {
	outFile, _ := os.OpenFile("dados.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer outFile.Close()

	// buffer := &bytes.Buffer{}
	// buffer.WriteString("41\n2\n35\n4\n12\n7\n19\n")

	// ResponseWriter := io.MultiWriter(buffer, outFile)
	// linkedlist := lista.New()
	// processData("Buffer + fila", linkedlist, buffer, ResponseWriter)
	// fmt.Println("buffer\n" + buffer.String())

	//HTTP
	http.HandleFunc("/processData", func(write http.ResponseWriter, req *http.Request) {
		struct_data := lista.New()

		ResponseWriter := io.MultiWriter(write, outFile)
		processData("HTTP + lista ligada", struct_data, req.Body, ResponseWriter)
	})

	fmt.Println("Servidor HTTP rodando em :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
