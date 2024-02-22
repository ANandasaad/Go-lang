package main

import "fmt"

const (
	first = iota
	second
	third = 'a'
)

func main() {

	a := 10
	b := 20

	fmt.Println(test(a, b))
	fmt.Println(first, second, third)
	if rune('a') == rune(third) {
		fmt.Println(true)
	}

}

func test(a int, b int) (int, error) {

	return a + b, nil
}
