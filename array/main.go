package main

var meuarray [3]int
var meuslice = []int{9, 5, 8, 3, 32, 1, 67, 14}
var meumapa = map[string]int{"henry": 100, "joão": 30, "josé": 80}

func main() {
	meuarray[0] = 20
	meuarray[1] = 21
	meuarray[2] = 19
	for nome, valor := range meumapa {

		print("O nome é ", nome)
		println(" e o salário é ", valor)
	}

}
