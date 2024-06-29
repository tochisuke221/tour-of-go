package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Board is a 9x9 array of integers
//
// 0: 見入力
// 1-9: 1-9の数字

type Board [9][9]int

func pretty(b Board) string {
  var buf bytes.Buffer

  for i := 0; i < 9; i++ {
    if i%3 == 0 {
      buf.WriteString("+---+----+----+\n")
    }
    for j := 0; j < 9; j++ {
      if j%3 == 0 {
        buf.WriteString("|")
      }
      buf.WriteString(strconv.Itoa(b[i][j])) 
    }
    buf.WriteString("|\n")
  }

  buf.WriteString("+---+----+----+\n")

  return buf.String()
}


func 
duplicated(c [10]int) bool {
  for k, v := range c {
    if k == 0 {
      continue
    }
    if v > 1 {
      return true
    }
  }
  return false
}

func verify(b Board) bool {
  // 行チェック
  for i := 0; i < 9; i++ {
    // 出現回数
    var c [10] int
    for j := 0; j < 9; j++ {
      c[b[i][j]]++
    }
    if 
    duplicated(c) {
      return false
    }
  }
  return true
}


func main() {
 b := Board{
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
  {0, 0, 0, 0, 0, 0, 0, 0, 0},
 }

 fmt.Println("%+v\n", pretty(b))
}