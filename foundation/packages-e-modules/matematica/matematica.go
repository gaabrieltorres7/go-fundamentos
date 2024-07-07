package matematica

// se eu usar a primeira letra maiúscula,
// a função será exportada e o contrário não.
// ou seja, não tenho acesso fora do pacote
// funciona tipo uma export do JS
func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro andando"
}
