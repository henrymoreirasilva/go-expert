package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func soma(a, b int) int {
	return a + b
}

func main() {
	templates := []string{
		"topo.html",
		"template.html",
		"rodape.html",
	}
	T := template.New("template.html")
	T.Funcs(template.FuncMap{"somar": soma})
	T.ParseFiles(templates...)

	err := T.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 45},
		{"C++", 30},
	})

	if err != nil {
		panic(err)
	}

}
