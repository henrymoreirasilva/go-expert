package main

type Endereco struct {
	Logradouro string
	Bairro     string
}

type Cliente struct {
	Codigo int
	Nome   string
	Endereco
}

func main() {
	henry := Cliente{
		Codigo: 1,
		Nome:   "Henry",
	}

	henry.Endereco.Logradouro = "rua x"

	print("Nome ")
	println(henry.Logradouro)
}
