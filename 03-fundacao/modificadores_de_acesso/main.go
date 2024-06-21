package main

import "mod_acc/matematica"

func main() {
	println(conexao())

	cc := matematica.Conta{
		Agencia: "4745",
		Conta:   "01002",
	}

	println(cc.Agencia)
}
