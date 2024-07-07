package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

/*com structs, não consigo trabalhar com herança,
  mas consigo trabalhar com composição
*/

func main() {

	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 24,
		Ativo: true,
	}

	gabriel.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", gabriel.Nome, gabriel.Idade, gabriel.Ativo)

}
