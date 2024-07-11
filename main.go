package main

import (
	"fmt"
	"time"
)

func main() {
	n := 1
	// nが競合するパターン
	go func() {
		for i := 2; i <= 5; i++ {
			fmt.Println(n, "*", i)
			n *= i
			time.Sleep(100)
		}
	}()
	
	for i := 1; i <= 10; i++{
		fmt.Println(n, "+", i)
		n += i
		time.Sleep(100)
	}
}
