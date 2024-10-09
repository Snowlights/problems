package main

func countPairs(nums []int, target int) int {
	ans := 0
	for i, v := range nums {
		for _, vv := range nums[i+1:] {
			if vv+v < target {
				ans++
			}
		}
	}
	return ans
}
