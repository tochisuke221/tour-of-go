package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// Board is a 9x9 array of integers
//
// 0: 未入力
// 1-9: 1-9の数字

type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer

	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			if b[i][j] == 0 {
				buf.WriteString(" ")
			} else {
				buf.WriteString(strconv.Itoa(b[i][j]))
			}
		}
		buf.WriteString("|\n")
	}

	buf.WriteString("+---+---+---+\n")

	return buf.String()
}

func duplicated(c [10]int) bool {
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
	// 列チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}

	// 行チェック
	for i := 0; i < 9; i++ {
		// 出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicated(c) {
			return false
		}
	}

	// 3x3 チェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}

	return true
}

func solved(b Board) bool {
	if !verify(b) {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// backtrack
// 参照渡しにする
func backtrack(b *Board) bool {
	if solved(*b) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// マスが０のとき
			if b[i][j] == 0 {
				for c := 1; c <= 9; c++ {  // 数字の候補を1から9までに変更
					b[i][j] = c
					// もし数字がルールに適合するなら
					if verify(*b) {
						// さらに探索する
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	b := Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if backtrack(&b) {
		fmt.Printf("解決した盤面:\n%s\n", pretty(b))
	} else {
		fmt.Println("解決策が見つかりませんでした")
	}
}
