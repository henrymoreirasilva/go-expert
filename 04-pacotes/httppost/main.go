package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonBuffer := bytes.NewBuffer([]byte(`{"name": "henry"}`))
	resp, err := c.Post("http://google.com", "application/json", jsonBuffer)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
