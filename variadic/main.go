package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	test("hello", nums...)
}

func test(str string, nums ...int) {
	for _, i := range nums {
		fmt.Println(i)
	}
	fmt.Println(str)
}
