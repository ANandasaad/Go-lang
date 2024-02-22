package main

import "fmt"

func main() {
	hello := "hello world"
	str := `hello`
	fmt.Println(hello[0])
	fmt.Println("\n", string(hello[0]))
	var substring = hello[:5]
	fmt.Println(substring)
	fmt.Println(str)
}
