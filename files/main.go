package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files")
	content := "Hey, world"
	file, err := os.Create("./workFile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readFile("./workFile.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)

	}
	fmt.Println(string(databyte))

}
