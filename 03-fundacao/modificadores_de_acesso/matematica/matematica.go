package matematica

var numero int
var Taxa int = 5

func Soma(a, b int) int {
	return a + b
}

type Conta struct {
	Agencia string
	Conta   string
	saldo   float64
}

func multiplica(a, b int) int {
	return a * b
}
