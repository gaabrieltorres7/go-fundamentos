package main

type MyNumber int

// constraint
type Number interface {
	~int | ~float64
	// uso "~" para poder utilizar o MyNumber e o Go
	// reconhecer que também é um int/float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// comparable constraint
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func main() {
	m := map[string]int{"Gabriel": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Gabriel": 100.20, "João": 2000.3, "Maria": 300.0}

	m3 := map[string]MyNumber{"Gabriel": 1000, "João": 2000, "Maria": 3000}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
