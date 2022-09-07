package _00_700

import "sort"

// 611
func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}
