package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		println(err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return
	}
	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}
