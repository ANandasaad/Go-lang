package main

import "fmt"

func main() {
	var username string = "Anand"
	fmt.Println(username)
	fmt.Printf("Variable is type of : %T\n", username)

	var isUserLoggedIn bool = false
	fmt.Println(isUserLoggedIn)
	fmt.Printf("Variable is type of : %T\n", isUserLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is type of : %T\n", smallValue)
}
