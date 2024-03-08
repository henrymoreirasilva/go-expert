package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("POST", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("accept", "application/json")
	resp, err := c.Do(req)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
