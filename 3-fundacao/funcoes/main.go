package main

type Cliente struct {
	Codigo int
	Nome   string
}

var Clientes [3]Cliente

func main() {

	Clientes[0].Codigo = 1
	Clientes[0].Nome = "Henry"

	println(Clientes[0].Nome)
}
