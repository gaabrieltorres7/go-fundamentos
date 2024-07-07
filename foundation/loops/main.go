package main

func main() {
	//for simples
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	//for com range
	for k, v := range numeros {
		println(k, v)
	}

	//for estilo while no go
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinito
	for {
		println("Hello, World!")
	}
}
