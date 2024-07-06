package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543))
}

// ...int = variadic parameter
// variadic parameter = aceita um número variável de argumentos
// parecido com o operador spread do javascript
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
