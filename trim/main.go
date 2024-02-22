package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "@hello@"
	fmt.Println(strings.Trim(str, "@"))
	fmt.Println(strings.TrimLeft(str, "@"))
	fmt.Println(strings.Split(str, ""))

}
