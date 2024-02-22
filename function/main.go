package main

import "fmt"

func main() {

	a := 10
	b := 20

	c := test(a, b)
	fmt.Println(c)
	d, f := swap("hello", "world")
	fmt.Println(d, f)
	fmt.Println(split(17))
}

func test(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return a, b
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x

	return x, y
}
