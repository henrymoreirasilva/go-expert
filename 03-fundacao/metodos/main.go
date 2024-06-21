package main

type Endereco struct {
	Logradouro string
	Bairro     string
}

type Cliente struct {
	Codigo int
	Nome   string
	Ativo  bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	henry := Cliente{
		Codigo: 1,
		Nome:   "Henry",
	}

	henry.Desativar()

	print("Ativo ")
	println(henry.Ativo)

}
