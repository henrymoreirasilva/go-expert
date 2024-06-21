package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {

	// struct para json - Marshal
	conta1 := Conta{Numero: 1, Saldo: 10}

	res, err := json.Marshal(conta1)
	if err != nil {
		panic(err)
	}
	print("Ex. 1: ")
	println(string(res))

	// struct para json - NewEncoder
	print("Ex. 2: ")
	err = json.NewEncoder(os.Stdout).Encode(conta1)
	if err != nil {
		panic(err)
	}

	// String par astruct - Unmarshal - array
	var contas []Conta
	jsonString := []byte(`[{"n": 3, "s": 19}, {"n": 4, "s": 9}, {"n": 5, "s": 32}]`)

	print("Json: ")
	println(string(jsonString))

	err = json.Unmarshal(jsonString, &contas)

	print("Ex. 3: ")
	println(contas[1].Saldo)

}