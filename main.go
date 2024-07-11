package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	var ch2 chan string
	go func() { ch1 <- 100 }()
	go func() { ch2 <- "hi" }()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}
