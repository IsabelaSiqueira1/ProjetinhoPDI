package main

import "fmt"

// type <nome da estrutura>	// struct {
// 	<nome do campo> <tipo do campo>
type Pessoa struct {
	Nome  string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Isabela", 24})            // Instanciando a struct pessoa com os valores Nome e Idade
	fmt.Println(Pessoa{Nome: "Bento", Idade: 24}) // Instanciando a struct pessoa com os valores Nome e Idade
	fmt.Println(Pessoa{Nome: "Isabela"})          //criando a struct pessoa com o valor Nome, mas sem o valor Idade, que será 0 por padrão, criara um zerovalue

	p1 := Pessoa{Nome: "Mkhael", Idade: 24} // Instanciando a struct pessoa com os valores Nome e Idade
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade) // Acessando os campos da struct pessoa

	p1.Idade = 25         // Atualizando o valor do campo Idade da struct pessoa
	fmt.Println(p1.Idade) // Imprimindo o novo valor do campo Idade da struct pessoa

	p2 := Pessoa{Nome: "Patrick", Idade: 16}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2) // Criando um slice de structs Pessoa e adicionando as structs p1 e p2
	fmt.Println(pessoas)              // Imprimindo o slice de structs Pessoa

	//structs + map
	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	// var alunos = map[string][]Pessoa{
	// 	"programação": {{Nome: "Isabela", Idade: 24}, {Nome: "Bento", Idade: 24}},
	// 	"Engenharia":  {{Nome: "Mkhael", Idade: 24}, {Nome: "Patrick", Idade: 16}},
	// }
	// fmt.Println(alunos) // Imprimindo o map de structs Pessoa

	prof := Profissao{p2, "dev"}
	fmt.Println(prof)              // Imprimindo a struct Profissao que contém a struct Pessoa e o campo Tipo
	fmt.Println(prof.Pessoa.Nome)  // Acessando o campo Nome da struct Pessoa dentro da struct Profissao
	fmt.Println(prof.Pessoa.Idade) // Acessando o campo Idade da struct Pessoa dentro da struct Profissao
	fmt.Println(prof.Tipo)         // Acessando o campo Tipo da struct Profissao
}
