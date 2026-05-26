package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// z := x
	for i := 0; i < 10; i++ {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Current z = %v (%v回目)\n", z, i+1)
		if diff := math.Abs(oldZ - z); diff < 0.0000000001 {
			break
		}
	}
	return z
}

func main() {
	x := 2.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
