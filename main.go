package main

import (
	"fmt"
	"os"
)

// インターフェースの定義
type Stringer interface{
	String() string
}

func ToStringer(v interface {}) (Stringer, error) {
	if s, ok := v.(Stringer); ok {
		return s, nil
	}

	return nil, MyError("キャストエラー")
}

type MyError string

func (e MyError) Error() string {
	return string(e)
}


type S string

func (s S) String() string {
	return string(s)
}


func main() {
	var v string = "hoge"

	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "エラーです:", err)
	}else {
		fmt.Println(s.String())
	}
}