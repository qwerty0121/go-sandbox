package main

import "fmt"

/**
 * 与えられた整数のスライスの要素の並び順を反転させる
 */
func Reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		a, b := nums[i], nums[len(nums)-1-i]
		nums[i], nums[len(nums)-1-i] = b, a
	}
}

func main() {
	// ケース1
	nums1 := []int{1, 2, 3, 4, 5}
	Reverse(nums1)
	fmt.Println(nums1) // [5, 4, 3, 2, 1]

	// ケース2
	nums2 := []int{10, 20}
	Reverse(nums2)
	fmt.Println(nums2) // [20, 10]

	// ケース3
	nums3 := []int{7}
	Reverse(nums3)
	fmt.Println(nums3) // [7]
}
