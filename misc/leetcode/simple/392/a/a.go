package main

func longestMonotonicSubarray(nums []int) int {
	pre, last, ans := 1, 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			pre++
		} else {
			pre = 1
		}
		if nums[i] < nums[i-1] {
			last++
		} else {
			last = 1
		}
		ans = max(ans, pre, last)
	}
	return ans
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
