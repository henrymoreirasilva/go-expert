package main

import "net/http"

func main() {
	blog := Blog{title: "Bem vindo ao blog"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog)

	http.ListenAndServe(":80", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home\n"))
}

type Blog struct {
	title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(b.title + "\n"))
}
