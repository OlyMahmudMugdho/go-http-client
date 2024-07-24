package client

import (
	"io"
	"net/http"
	"net/url"
)

func GetRequest(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	resBody, readError := io.ReadAll(response.Body)

	if readError != nil {
		return nil, readError
	}

	return resBody, nil

}

func PostFormData(url string, data url.Values) ([]byte, error) {
	response, error := http.PostForm(url, data)
	if error != nil {
		return nil, error
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return responseData, nil
}
