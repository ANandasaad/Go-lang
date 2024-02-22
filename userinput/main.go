package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the Go programming language"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the pizza")

	// comma ok // error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating!", input)
	fmt.Printf("Variable is type of : %T\n", input)
}
