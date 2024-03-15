package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, par := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + par + "/json/")

		if err != nil {
			panic("Erro ao acessar o serviço")
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			panic("Erro ao ler conteúdo")
		}

		var Dados ViaCep

		err = json.Unmarshal(res, &Dados)
		if err != nil {
			panic("Erro ao converter json")
		}

		println(Dados.Logradouro)
	}

}
