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
	EncodeJson()
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
