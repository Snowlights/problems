package main

func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx, s := 0, 0
	for i, x := range nums {
		mx = max(mx, x) // 前缀最大值
		s += x + mx     // 累加前缀的得分
		ans[i] = int64(s)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
