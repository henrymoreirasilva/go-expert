package main

type Pessoa interface {
	Desativar()
	Ativar()
}
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

type Empresa struct {
	Codigo   int
	Fantasia string
	Ativo    bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	println(c.Ativo)
}
func (c Empresa) Desativar() {
	c.Ativo = false
	println(c.Ativo)
}

func (c Cliente) Ativar() {
	c.Ativo = true
	println(c.Ativo)
}

func (c Empresa) Ativar() {
	c.Ativo = true
	println(c.Ativo)
}

func main() {
	// henry := Cliente{
	// 	Codigo: 1,
	// 	Nome:   "Henry",
	// }

	zoomwi := Empresa{
		Codigo:   1,
		Fantasia: "zoom prop",
	}

	zoomwi.Desativar()

}
