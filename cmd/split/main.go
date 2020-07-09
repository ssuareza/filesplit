package main

import (
	"fmt"
	"log"

	"github.com/ssuareza/filesplit"
)

func main() {
	chunks, err := filesplit.Split("/tmp/file.dat")
	if err != nil {
		log.Fatal(err)
	}
	for _, chunk := range chunks {
		fmt.Println(chunk.Name)
	}
}
