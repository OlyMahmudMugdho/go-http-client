package main

import (
	"fmt"

	"github.com/OlyMahmudMugdho/go-http-client/client"
)

func main() {
	getUrl := "http://localhost:8080"
	postUrl := getUrl + "/post"

	/* data, _ := client.GetRequest(getUrl)
	fmt.Println(string(data))

	form := url.Values{}
	form.Add("name", "mugdho")
	formData, _ := client.PostFormData(postUrl, form)
	fmt.Println(string(formData)) */

	response, _ := client.PostRequest(postUrl, "application/json", `{"name" :"mugdho"}`)
	fmt.Println(string(response))

}
