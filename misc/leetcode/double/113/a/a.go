package main

import "sort"

func minimumRightShifts(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	for i := 1; i <= len(nums); i++ {
		nums = append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...)
		if sort.IntsAreSorted(nums) {
			return i
		}
	}

	return -1
}
