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
	T := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := T.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 45},
		{"C++", 30},
	})

	if err != nil {
		panic(err)
	}

}
