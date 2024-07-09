package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
	// json:"n" é um tag. é um metadado que diz
	// para o pacote json como serializar e deserializar
	// mesmo que o nome da struct seja diferente
}

func main() {
	// Marshal retorna o JSON da struct
	// guardo o retorno em uma variável do tipo []byte
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Encoder
	//pego o valor, serializo e escrevo no stdout
	err2 := json.NewEncoder(os.Stdout).Encode(conta)
	if err2 != nil {
		println(err2)
	}

	// Unmarshal
	exampleJSON := []byte(`{"n":1,"s":100}`)
	var contaX Conta // variável em branco
	// deserializo o JSON e guardo na variável em branco
	// o segundo parâmetro é um ponteiro para a variável
	// representa o endereço de memória da variável
	err3 := json.Unmarshal(exampleJSON, &contaX)
	if err3 != nil {
		println(err3)
	}
	println(contaX.Saldo)
}
