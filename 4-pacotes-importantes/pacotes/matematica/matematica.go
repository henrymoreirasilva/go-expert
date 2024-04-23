package matematica

// A função deve ser declarada com a primeira maiúscula para que seja exportada para
// os outros módulos. Com letra inicial minúscula, o escopo é local do módulo.
// O mesmo vale para variáveis, structs....
// Inclusive para propriedades de structs.
func Soma[T int | float64](a, b T) T {
	return a + b
}
