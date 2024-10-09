package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 2; i < len(nums); i += 3 {
		if nums[i]-nums[i-2] <= k {
			ans = append(ans, nums[i-2:i+1])
		}
	}
	if len(ans)*3 != len(nums) {
		return nil
	}
	return ans
}
