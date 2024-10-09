package main

func maxSubarrayLength(nums []int, k int) int {
	h := make(map[int]int)
	ans, l := 0, 0
	for r, v := range nums {
		h[v]++
		for h[v] > k {
			h[nums[l]]--
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
