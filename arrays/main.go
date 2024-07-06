package main

import "fmt"

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
	var arr [3]int //tamanho fixo
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(len(arr) - 1)    // 2
	fmt.Println(arr)             // [1 2 3]
	fmt.Println(len(arr))        // 3
	fmt.Println(arr[len(arr)-1]) // 3 -> take the last element of the array

	for i, v := range arr {
		fmt.Printf("O valor do índice é %d e o valor é %d \n", i, v) // 0 1, 1 2, 2 3
	}

	//string %v e numero %d

}
