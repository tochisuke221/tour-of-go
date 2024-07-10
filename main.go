package main

import (
	"fmt"
)

// インターフェースの定義
type Stringer interface {
	String() string
}


type MyString string

func (s MyString) String() string {
	return "これはMyString型です"
}

type MyInt int
func (i MyInt) String() string {
	return "これはMyInt型です"
}

type MyBool bool
func(b MyBool) String() string {
	return "これはMyBool型です"
}


func F(s Stringer){
	switch v := s.(type) {
		case MyString:
			fmt.Println(string(v) , "S")
		case MyInt:
			fmt.Println(int(v) , "I")
		case MyBool:
			fmt.Println(bool(v) , "B")
	}
}

type Hoge struct {
	N int
}

type Fuga struct {
	Hoge
}



func main() {
	var s Stringer
	s = MyString("文字列です")
	F(s)
	s = MyInt(10)
	F(s)
	s = MyBool(true)
	F(s)

	f := Fuga{Hoge{100}}

	fmt.Println(f.N)
}