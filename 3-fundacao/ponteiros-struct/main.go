package main

type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{}
}

func (c Conta) SomaSaldo(valor int) int {
	c.Saldo = c.Saldo + valor
	return c.Saldo
}

func main() {
	minhaConta := NewConta()

	println(minhaConta.SomaSaldo(10))
	println(minhaConta.Saldo)
}
