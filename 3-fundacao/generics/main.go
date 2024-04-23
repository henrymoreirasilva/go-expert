package main

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main() {
	m := map[string]int{"henry": 10, "gabriel": 20, "andreza": 30}
	m2 := map[string]float64{"henry": 10.33, "gabriel": 20, "andreza": 30}
	println(Soma(m))
	println(Soma(m2))
}
