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

func main() {
	templates := []string{
		"topo.html",
		"template.html",
		"rodape.html",
	}
	T := template.Must(template.New("template.html").ParseFiles(templates...))

	err := T.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 45},
		{"C++", 30},
	})

	if err != nil {
		panic(err)
	}

}
