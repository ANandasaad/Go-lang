package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	result, _ := url.Parse(s)
	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)
	fmt.Println(result.Fragment)
	fmt.Println(result.Port())
	params := result.Query()
	fmt.Println(params["k"])

	for _, v := range params {
		fmt.Println(v)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=anand",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
