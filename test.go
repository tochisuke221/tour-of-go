package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))


	var i, j int = 1, 3
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
