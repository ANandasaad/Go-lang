package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(2024, time.July, 17, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("2006-01-02 15:04:05"))
}
