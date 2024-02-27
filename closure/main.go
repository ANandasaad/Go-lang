package main

import "fmt"

func main() {
	x := test()
	fmt.Println(x())
	fmt.Println(x())

}

func test() func() string {
	a := "hello"
	return func() string {

		return a
	}

}
