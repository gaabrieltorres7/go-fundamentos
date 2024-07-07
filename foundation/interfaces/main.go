package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// a interface é implementada implicitamente
// por qualquer tipo que tenha o método Desativar
type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	// gabriel := Cliente{
	// 	Nome:"Gabriel",
	// 	Idade: 24,
	// 	Ativo: true,
	// }
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
}
