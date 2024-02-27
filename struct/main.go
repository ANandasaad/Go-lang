package main

import "fmt"

type Vertex struct {
	X int
	Y string
}

func main() {
	v := Vertex{1, "world"}
	fmt.Println(Vertex{1, "hello"})
	fmt.Println(v.X)
	fmt.Println(v.Y)
	p := &v
	fmt.Println(p.Y)
}
