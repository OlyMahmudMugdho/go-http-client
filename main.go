package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080"

	data, _ := GetRequest(url)

	fmt.Println(string(data))
}

func GetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	resBody, readError := io.ReadAll(response.Body)

	if readError != nil {
		fmt.Println(readError)
		return nil, readError
	}

	return resBody, nil

}
