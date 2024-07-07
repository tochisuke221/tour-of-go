package main

import (
	"fmt"
)
	
func main() {
	for i :=0 ; i < 10; i++ {
		hantei(&i)
  }

	n, m := 10, 20
	swap2(&n, &m)
	// println(n, m)
	
	println(n, m)



}

func hantei(i *int) {
	if *i%2 == 0 {
		fmt.Printf("%d-偶数\n", *i)
	} else {
		fmt.Printf("%d-奇数\n", *i)
	}
}

func swap(a, b int)(int, int){
	return b, a
}

func swap2(a, b *int)(int, int){
	var tmp = *a
	*a = *b
	*b = tmp

	return *a, *b
}