package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", BuscaCep)
	http.ListenAndServe(":80", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var paramCep string
	paramCep = r.URL.Query().Get("cep")
	if paramCep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(paramCep))
}
