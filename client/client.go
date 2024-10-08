package client

import (
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
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

	response.Body.Close()

	return resBody, nil

}

func PostRequest(url string, contentType string, data string) ([]byte, error) {
	response, error := http.Post(url, contentType, strings.NewReader(data))
	if error != nil {
		return nil, error
	}
	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	response.Body.Close()

	return responseData, nil
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

	response.Body.Close()

	return responseData, nil
}

func ReverseProxy(w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse("http://localhost:8081")
	proxy := httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(url)
			pr.Out.Host = pr.In.Host
		},
	}
	proxy.ServeHTTP(w, r)
}
