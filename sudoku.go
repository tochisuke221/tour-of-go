package main

import (
	"fmt"
	"math/rand"
	"time"
	"slices"
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



	t := time.Now().UnixNano()
	rand.Seed(t)
	if result := rand.Intn(6); result%6 == 0 {
		fmt.Println("大吉")
	} else if result%5 == 0 || result%4 == 0{
		fmt.Println("中吉")
	} else if result%3 == 0 || result%2 == 0{
		fmt.Println("吉")
	} else{
		fmt.Println("凶")
	}

	var f float64 = 10
	var n int = int(f)
	println(n)


	sum := 5 + 6 + 3

	avg := float64(sum / 3)

	if avg > 2.5 {
		println("good")
	}else {
		println("bad")
	}

	// var p struct {
	// 	name string
	// 	age int
	// }


	// p := struct {
	// 	name string
	// 	age int
	// }{
	// 	name: "gopher",
	// 	age: 10,
	// }

	// println(p.name, p.age)

	// ns3 := [5]int{10, 20, 30, 40, 50}

	// println(ns3[1:3])

	// ns := make([]int, 5, 10)
	// println(ns)
	// println(len(ns))
	// println(cap(ns))

	// ns = append(ns, 10, 10)
	// println(ns)
	// println(len(ns))
	// println(cap(ns))

	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4

	fmt.Println(ns[n:])
	fmt.Println(ns[:m])

	ms := ns[:m:5]

	fmt.Println(cap(ms))

	ns = slices.Delete(ns, 1, 3)

	fmt.Println(ns)

	ns = slices.Insert(ns, 1, 100, 200)

	fmt.Println(ns)

	fmt.Println(slices.Contains(ns, 100))

	slices.Sort(ns)
	fmt.Println(ns)







		
}