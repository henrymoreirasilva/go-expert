package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// pacotes os e io
	arquivo, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	t, err := io.WriteString(arquivo, "Bom dia")
	if err != nil {
		panic(err)
	}
	println(t)

	// pacote bufio
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))

	}
	arquivo2.Close()

}
