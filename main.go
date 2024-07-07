package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "メッセージ")
var n int

func init() {
	// ポインタを指定して設定を予約
	flag.IntVar(&n, "n", 1, "回数")
}


func main() {
  fmt.Println(os.Args)

  flag.Parse()

  fmt.Println(strings.Repeat(*msg, n))
}

