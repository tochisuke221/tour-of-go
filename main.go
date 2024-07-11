package main

import (
	"fmt"
	"time"
)

func main() {
	var v int
	go func() {// ゴールーチン-1
		time.Sleep(1 * time.Second)
		v = 100
	}()
	go func() {// ゴールーチン-2
		time.Sleep(1 * time.Second)
		fmt.Println(v)
	}()
	time.Sleep(2 * time.Second)
}
