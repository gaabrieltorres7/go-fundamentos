package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
	Ativo      bool
}

// utilizo composição de structs
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

//acesso de duas formas diferentes
//gabriel.Endereco.Logradouro
//gabriel.Logradouro

// método associado a struct Cliente
// (c Cliente) é onde eu defino quem é o receptor do método
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 24,
		Ativo: true,
	}

	gabriel.Desativar()
}
