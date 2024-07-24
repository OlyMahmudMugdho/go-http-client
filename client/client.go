package client

import (
	"fmt"
	"io"
	"net/http"
)

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
