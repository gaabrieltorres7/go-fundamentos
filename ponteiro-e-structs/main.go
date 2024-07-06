package main

type Conta struct {
	saldo int
}

// crio uma funcao que retorna o endereco da memoria
// de uma nova Conta
// espécie de construtor
func NewConta() *Conta {
	return &Conta{saldo: 0} // &retorna o endereco da memoria
}

// *Conta é um ponteiro para Conta com isso eu altero
// o valor de saldo
func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
