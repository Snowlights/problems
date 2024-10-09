package main

func minLengthAfterRemovals(nums []int) int {
	cnt := make(map[int]int)
	for _, v := range nums {
		cnt[v]++
	}
	mx := 0
	for _, v := range cnt {
		mx = max(mx, v)
	}

	if mx*2 > len(nums) {
		return 2*mx - len(nums)
	}
	return len(nums) % 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
