//go:generate stringer -type=Fruit

package main

import (
	"fmt"
)

type F struct{
	Name string
	Age int
}


func (f *F) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "John",
		Age: 25,
	}

	fmt.Printf("%+v\n", f)
}

type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)