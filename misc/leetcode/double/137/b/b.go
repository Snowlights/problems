package main

func resultsArray(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	for i := range ans {
		ans[i] = -1
	}
	cnt := 0
	for i, x := range nums {
		if i == 0 || x == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= k {
			ans[i-k+1] = x
		}
	}
	return ans
}
