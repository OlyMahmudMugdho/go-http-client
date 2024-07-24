package main

import (
	"fmt"
	"net/url"

	"github.com/OlyMahmudMugdho/go-http-client/client"
)

func main() {
	myUrl := "http://localhost:8080"

	data, _ := client.GetRequest(myUrl)
	fmt.Println(string(data))

	form := url.Values{}
	form.Add("name", "mugdho")
	formUrl := myUrl + "/post"
	formData, _ := client.PostFormData(formUrl, form)
	fmt.Println(string(formData))
}
