package main

import "fmt"

// Reverse は、与えられた整数のスライスの要素の並び順を反転させます。
// ※slices.Reverse() と同様の動作をします。
func Reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
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

	// 以下、エッジケース

	// ケース4
	nums4 := []int{}
	Reverse(nums4)
	fmt.Println(nums4) // []
}
