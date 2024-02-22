package main

import (
	"fmt"
	"strings"
)

func main() {
	// sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// for sum < 100 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	str := "hello world"
	fmt.Println("Print", str)
	for i := 0; i < len(str); i++ {

		fmt.Print(string(str[i]))
	}

	runes := []rune(str)
	fmt.Println(runes)

	for _, r := range runes {
		fmt.Printf("%c ", r)
	}

	slipt := strings.Split(str, "")
	fmt.Println(slipt)
	for i := 0; i < len(slipt); i++ {
		fmt.Println(slipt[i])
	}

}
