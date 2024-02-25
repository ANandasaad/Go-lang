package main

import "fmt"

type data int

func main() {
	value1 := data(23)
	value2 := data(24)

	res := value1.multiply(value2)
	fmt.Println(res)
}

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}
