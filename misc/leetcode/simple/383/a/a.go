package main

func returnToBoundaryCount(nums []int) int {
	ans, cur := 0, 0
	for _, v := range nums {
		cur += v
		if cur == 0 {
			ans++
		}
	}
	return ans
}
