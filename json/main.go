package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Platform string   `json:"website"`
	Price    int      `json:"price"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourse := []course{
		{"React js", "react.com", 100, "123abc", []string{"react", "node"}},
		{"React js", "react.com", 100, "123abc", nil},
	}

	finalJson, err := json.MarshalIndent(lcocourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDecode := []byte(`
	  {
		"name": "React js",
		"website": "react.com",
		"price": 100,
		"tags": ["react","node"]
	}
	`)

	var lcoCourse course

	checkJson := json.Valid(jsonDecode)
	if checkJson {
		json.Unmarshal(jsonDecode, &lcoCourse)
		fmt.Println(lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID")
	}

}
