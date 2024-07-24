package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/go-http-client/client"
)

func main() {
	url := "http://localhost:8080"

	data, _ := client.GetRequest(url)

	fmt.Println(string(data))
}
