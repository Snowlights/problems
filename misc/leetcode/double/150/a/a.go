package main

func sumOfGoodNumbers(nums []int, k int) (ans int) {
	for i, x := range nums {
		if (i < k || x > nums[i-k]) && (i+k >= len(nums) || x > nums[i+k]) {
			ans += x
		}
	}
	return
}
