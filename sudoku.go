package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const message = "Hello, World!"

	println(message)

	// const (
	// 	a = 1 + 2
	// 	b
	// 	c
	// 	// The number of rows and columns in the grid
	// )

	//  println(a, b, c)


	 const (
		Draft = iota
		Open
		Pending
	 )
	 
	 println(Draft, Open, Pending)

	 x := 3
	 if x == 1 {
		println("x is 1")
	 }else if x == 2 {
		println("x is 2")
	}else{
			println("x is not 1 or 2")
	}

	if z := 5; z < 10 {
		println("z is less than 10")
	}

	var test []int= []int{1, 2}

	println(test)

	for i, v := range []int{1,2,3} {
		println(i, v)
	}

	for i := 0; i< 10; i++ {
		println(i)
	}

	// for i := 1; i < 101; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%d-偶数\n", i)
	// 	}else{
	// 		fmt.Printf("%d-奇数\n", i)
	// 	}
	// }




	if result := rand.Intn(6); result%6 == 0 {
		fmt.Println("大吉")
	} else if result%5 == 0 || result%4 == 0{
		fmt.Println("中吉")
	} else if result%3 == 0 || result%2 == 0{
		fmt.Println("吉")
	} else{
		fmt.Println("凶")
	}

		
}