package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// getResponse()
	// postRequest()
	postFormData()
}

func postFormData() {
	const myUrl = "http://localhost:3003/api/v1/city/create-city"

	data := url.Values{}
	data.Add("name", "Duai")
	data.Add("popularLocalityName", "Dubai")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func getResponse() {
	const url = "http://localhost:3003/api/v1/hotel/get-hotels"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("response length:", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func postRequest() {

	const url = "http://localhost:3003/api/v1/city/create-city"

	requestBody := strings.NewReader(`{
		"name":"Mumbai",
		"popularLocalityName":"navi mumbai
	}`)
	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	cotent, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(cotent))

}
