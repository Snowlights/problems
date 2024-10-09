package main

func minOperations(nums []int, k int) int {
	ans := 0
	for _, v := range nums {
		if v < k {
			ans++
		}
	}
	return ans
}
