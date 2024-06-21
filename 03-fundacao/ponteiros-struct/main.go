package main

var minhaConta = NewConta()

type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{}
}

func (c *Conta) SomaSaldo(valor int) int {
	c.Saldo = c.Saldo + valor
	return c.Saldo
}

func Teste() {
	minhaConta.Saldo = 22
}

func main() {

	Teste()

	println(minhaConta.SomaSaldo(10))
	println(minhaConta.Saldo)
}
