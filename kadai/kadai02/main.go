// 課題2: 文字列の出現頻度カウント

package main

import (
	"fmt"
	"strings"
)

// CountWords は、与えられた文字列（英文）に含まれる単語の出現回数をカウントし、マップで返す関数です。
func CountWords(s string) map[string]int {
	// 文字列をスペースで分割して単語のスライスを作成
	words := strings.Fields(s)
	if len(words) == 0 {
		return make(map[string]int) // 空のマップを返す
	}

	// 単語の出現回数をカウントするためのマップを作成
	// ※マップの初期容量を単語の数に設定して、メモリの再割り当てを防止
	m := make(map[string]int, len(words))

	// 単語のスライスをループして、マップに出現回数をカウント
	for _, word := range words {
		m[word]++
	}

	return m
}

func main() {
	// ケース1
	text1 := "apple banana apple"
	counts1 := CountWords(text1)
	fmt.Println(counts1) // map[string]int{"apple": 2, "banana": 1}

	// ケース2
	text2 := "go go go"
	counts2 := CountWords(text2)
	fmt.Println(counts2) // map[string]int{"go": 3}

	// ケース3
	text3 := ""
	counts3 := CountWords(text3)
	fmt.Println(counts3) // map[string]int{}
}
