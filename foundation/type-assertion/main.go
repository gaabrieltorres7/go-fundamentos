package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Gabriel Torres"
	//type assertion
	println(minhaVar.(string))
	res, ok := minhaVar.(int) // res = 0, ok = false
	// res = 0 porque o valor de minhaVar não é um int
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2) // panic: interface conversion: interface {} is string, not int
}
