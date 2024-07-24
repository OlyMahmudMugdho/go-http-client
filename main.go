package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:8080"
	response, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
		return
	}
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
