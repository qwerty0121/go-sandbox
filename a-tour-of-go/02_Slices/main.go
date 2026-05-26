package main

import "golang.org/x/tour/pic"

// import "fmt"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		p[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// dot := uint8((x * y) / 2)
			// dot := uint8(x * y)
			dot := uint8(x ^ y)
			p[y][x] = dot
		}
	}
	return p
}

func main() {
	pic.Show(Pic)

	// fmt.Println(Pic(10, 20))
}
