package main

func countAlternatingSubarrays(nums []int) int64 {
	ans, cnt := 0, 0
	for i, v := range nums {
		if i > 0 && nums[i-1] != v {
			cnt++
		} else {
			cnt = 1
		}
		ans += cnt
	}
	return int64(ans)
}
