package main

import "fmt"

func main() {
	var a = 42
	var p *int = &a

	fmt.Println(a)
	fmt.Println("Memory address", &a)
	fmt.Println("Pointer address", *p)
	*p = 21
	fmt.Println("value of a", a)
}
