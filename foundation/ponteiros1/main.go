package main

func main() {
	// Memória -> Endreço -> Valor
	a := 10
	var ponteiro *int = &a // &a -> aponto para o endereço de memória e *int -> tipo do ponteiro
	*ponteiro = 20         // *ponteiro -> estou acessando o valor que está no endereço de memória
	b := &a                // b é um ponteiro que aponta para o endereço de memória de a
	*b = 30                // acesso o valor que está no endereco de memoria e consequentemente altero o valor de a
	println(a)
}
