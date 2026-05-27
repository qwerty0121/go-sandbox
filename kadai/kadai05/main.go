// 課題5: インターフェースの基本

package main

import (
	"fmt"
	"math"
)

// Shape は平面を表すインターフェースです。
type Shape interface {
	// Area は平面の面積を計算するメソッドです。
	Area() float64
}

// Rectangle は四角形を表す構造体です。
type Rectangle struct {
	// 横
	Width int
	// 縦
	Height int
}

// Area は四角形の面積を計算するメソッドです。
// 四角形の面積を返却します。
func (r *Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

// Circle は円を表す構造体です。
type Circle struct {
	// 半径
	Radius int
}

// Area は円の面積を計算するメソッドです。
// 円の面積を返却します。
func (r *Circle) Area() float64 {
	return math.Pow(float64(r.Radius), 2) * math.Pi
}

func main() {
	// ケース1
	r := Rectangle{Width: 3, Height: 4}
	fmt.Printf("ケース1: %v.Area() = %v\n", r, r.Area())

	// ケース2
	c := Circle{Radius: 2}
	fmt.Printf("ケース2: %v.Area() = %v\n", c, c.Area())

	// ケース3
	var sr, sc Shape = &r, &c
	fmt.Printf("ケース3(Rectangle): %v.Area() = %v\n", sr, sr.Area())
	fmt.Printf("ケース3(Circle): %v.Area() = %v\n", sc, sc.Area())
}
