package main

const hello = "Hello, World!"

var ( //variavel de escopo global
	b bool    = true
	c int     = 10
	d string  = "Gabriel"
	e float64 = 1.2
	// b bool false by default
	// c int 0 by default
	// d string "" by default
	// e float64 0.0 by default
)

func main() {
	// a := "Xablau" //forma mais facil de declarar uma variavel
	// a = "Xena" //atribui outro valor a essa variavel
	// var z int = 10 //se voce declarar a variavel e nao usar, o compilador nao deixa rodar
	println(b)
}
