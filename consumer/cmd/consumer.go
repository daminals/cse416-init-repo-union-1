package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// var (
// 	addr = flag.String("addr", "localhost:50051", "the address to connect to")
// )

func main() {
	res, err := http.Get("http://127.0.0.1:8080/file/1234")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseBody))
}
