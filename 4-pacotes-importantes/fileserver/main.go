package main

import "net/http"

func main() {
	fileServe := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServe)

	http.ListenAndServe(":80", mux)
}
