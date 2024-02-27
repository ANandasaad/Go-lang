package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}

	fmt.Println(arr[4])

	primes := [6]int{1, 2, 3, 4, 5}
	var a []int = primes[0:3] // slice
	fmt.Println(a)

	//slice is reference for arrays

	names := []string{"anand", "baz", "cancel", "order"}

	fmt.Println(names)

	var b []string = names[0:1]
	fmt.Println(b)
	var c []string = names[1:3]
	fmt.Println(c)
	c[1] = "product"
	fmt.Println(names)
	var d []string = names[1:]
	fmt.Println(d)
}
