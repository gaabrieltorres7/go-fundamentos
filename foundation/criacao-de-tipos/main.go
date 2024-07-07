package main

const hello = "Hello, World!"

type ID int // posso criar meu proprio tipo

var ( //variavel de escopo global
	b bool    = true
	c int     = 10
	d string  = "Gabriel"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	println(f)
}
