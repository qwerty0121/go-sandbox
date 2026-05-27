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
	testCases := []struct{ a, b int }{
		{10, 2},
		{5, 0},
		{0, 3},
	}

	for i, testCase := range testCases {
		fmt.Printf("--- ケース %v ---\n", i+1)
		result, err := Divide(testCase.a, testCase.b)
		fmt.Printf("%v / %v = %v, %v\n", testCase.a, testCase.b, result, err)
	}
}
