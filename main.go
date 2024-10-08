package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OlyMahmudMugdho/go-http-client/client"
	"github.com/OlyMahmudMugdho/go-http-client/utils"
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

	utils.GetArgs()

	router := http.NewServeMux()
	router.HandleFunc("/", client.ReverseProxy)
	fmt.Println("server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
