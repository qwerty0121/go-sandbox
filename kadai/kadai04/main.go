// 課題4: 明示的なエラーハンドリング

package main

import (
	"errors"
	"fmt"
)

// Divide は2つの整数を受け取り、割り算の結果を返す関数です。
// 引数 a, b には整数を指定います。
// b が0でない場合、a / bを返します。
// b が0である場合はエラーを返します。
func Divide(a, b int) (int, error) {
	if b == 0 {
		// ゼロ除算が発生した場合はエラー
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func main() {
	// ケース1
	a1, b1 := 10, 2
	result1, err := Divide(a1, b1)
	fmt.Printf("ケース1: %v / %v = %v, %v\n", a1, b1, result1, err)

	// ケース2
	a2, b2 := 5, 0
	result2, err := Divide(a2, b2)
	fmt.Printf("ケース2: %v / %v = %v, %v\n", a2, b2, result2, err)

	// ケース3
	a3, b3 := 0, 3
	result3, err := Divide(a3, b3)
	fmt.Printf("ケース3: %v / %v = %v, %v\n", a3, b3, result3, err)
}
